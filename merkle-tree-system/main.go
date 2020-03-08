package main

import (
	"./saver"
	"errors"
	"fmt"
	"github.com/LimeChain/merkletree"
	"github.com/LimeChain/merkletree/memory"
	"github.com/LimeChain/merkletree/postgres"
	merkleRestAPI "github.com/LimeChain/merkletree/restapi/baseapi"
	env "github.com/Netflix/go-env"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/simonleung8/flags"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Environment struct {
	Blockchain struct {
		NodeURL         *string `env:"NODE_URL"`
		Secret          *string `env:"SECRET_KEY"`
		ContractAddress *string `env:"CONTRACT_ADDRESS"`
		SavePeriod      int     `env:"SAVE_PERIOD"`
	}

	Database struct {
		Host    *string `env:"DB_HOST"`
		Port    int     `env:"DB_PORT"`
		User    *string `env:"DB_USER"`
		Pass    *string `env:"DB_PASS"`
		DBName  *string `env:"DB_NAME"`
		DBExtra *string `env:"DB_EXTRA"`
	}

	Port int `env:"API_PORT"`

	Extras env.EnvSet
}

var treeLen int

type savedResponse struct {
	merkleRestAPI.MerkleAPIResponse
	Length int `json:"lastSavedIndex"`
}

func getSaved(tree merkletree.ExternalMerkleTree) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if tree.Length() == 0 {
			render.JSON(w, r, savedResponse{merkleRestAPI.MerkleAPIResponse{true, ""}, -1})
			return
		}
		render.JSON(w, r, savedResponse{merkleRestAPI.MerkleAPIResponse{true, ""}, treeLen})
		return
	}
}

func createAndStartAPI(tree merkletree.ExternalMerkleTree, port int) {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.Compress(6, "gzip"),
		middleware.RedirectSlashes,
		middleware.Recoverer,
		middleware.NoCache,
	)

	router.Route("/v1", func(r chi.Router) {
		treeRouter := chi.NewRouter()
		treeRouter = merkleRestAPI.MerkleTreeStatus(treeRouter, tree)
		treeRouter = merkleRestAPI.MerkleTreeInsert(treeRouter, tree)
		treeRouter = merkleRestAPI.MerkleTreeHashes(treeRouter, tree)
		treeRouter = merkleRestAPI.MerkleTreeRawInsert(treeRouter, tree)

		treeRouter.Get("/saved", getSaved(tree))
		r.Mount("/api/merkletree", treeRouter)
	})

	fmt.Printf("Starting REST Api at port %v\n", port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), router))
}

func createSaver(tree merkletree.FullMerkleTree, nodeUrl, privateKeyHex, contractAddress string, periodSeconds int) {
	treeSaver, err := saver.NewSaver(
		nodeUrl,
		privateKeyHex,
		contractAddress,
		tree)
	if err != nil {
		panic(err)
	}

	go func() {
		len := 0
		timeout := time.Duration(periodSeconds) * time.Second
		for {
			tree.Recalculate()
			savedRoot, err := treeSaver.FetchRoot()
			if err != nil {
				fmt.Println("ERR: Could not save the tree root")
				fmt.Println(err.Error())
				time.Sleep(timeout)
				continue
			}

			if savedRoot == tree.Root() {
				fmt.Printf("Same root (%v) found in the chain. Skipping this iteration\n", savedRoot)
				treeLen = tree.Length()
				time.Sleep(timeout)
				continue
			}

			treeLen = tree.Length()
			if treeLen > len {
				fmt.Printf("Submitting new tree root to the chain (%v)\n", tree.Root())
				tx, err := treeSaver.TriggerSave()
				if err != nil {
					fmt.Println("ERR: Could not save the tree root")
					fmt.Println(err.Error())
				} else {
					fmt.Printf("Transaction (%v) mined\n", tx.TxHash.Hex())
					fmt.Printf("Gas used (%v)\n", tx.GasUsed)
					len = treeLen
				}
			} else {
				fmt.Println("No changes to submit. Skipping this iteration")
			}
			time.Sleep(timeout)
		}
	}()

	fmt.Printf("Started saver on %v seconds\n", periodSeconds)
	fmt.Printf("Node url %v\n", nodeUrl)
	fmt.Printf("Verifier contract address %v \n", contractAddress)

}

func setupFlags() flags.FlagContext {
	fc := flags.New()

	fc.NewBoolFlag("envirnonment-variables", "e", "Connection string for the postgres database")
	fc.NewStringFlag("database-connection", "db", "Connection string for the postgres database")
	fc.NewStringFlag("node-url", "n", "url to the ethereum node to connect")
	fc.NewStringFlag("secret", "s", "private key for the ethereum saver")
	fc.NewStringFlag("conntract-address", "c", "Address to the verifier contract")
	fc.NewIntFlagWithDefault("port", "ap", "port to run the API on", 8080)
	fc.NewIntFlagWithDefault("period", "p", "period to try and save the new root", 15)

	fc.Parse(os.Args...)

	return fc
}

func loadPostgreTree(connStr string) merkletree.FullMerkleTree {
	tree := postgres.LoadMerkleTree(memory.NewMerkleTree(), connStr)
	fmt.Printf("Merkle tree loaded. Length : %v, Root : %v\n", tree.Length(), tree.Root())
	return tree
}

func main() {
	fc := setupFlags()

	var connStr, nodeURL, privateKeyHex, contractAddress string
	var period, port int

	if fc.Bool("e") {
		var environment Environment
		_, err := env.UnmarshalFromEnviron(&environment)
		if err != nil {
			log.Fatal(err)
		}
		connStr = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s %s", *environment.Database.Host, environment.Database.Port, *environment.Database.User, *environment.Database.Pass, *environment.Database.DBName, *environment.Database.DBExtra)
		nodeURL = *environment.Blockchain.NodeURL
		privateKeyHex = *environment.Blockchain.Secret
		contractAddress = *environment.Blockchain.ContractAddress
		period = environment.Blockchain.SavePeriod
		port = environment.Port
	} else {
		if !fc.IsSet("db") {
			panic(errors.New("No db flag set"))
		}

		if !fc.IsSet("n") {
			panic(errors.New("No node-url flag set"))
		}

		if !fc.IsSet("s") {
			panic(errors.New("No secret flag set"))
		}

		if !fc.IsSet("c") {
			panic(errors.New("No conntract-address flag set"))
		}

		connStr = fc.String("db")
		nodeURL = fc.String("n")
		privateKeyHex = fc.String("s")
		contractAddress = fc.String("c")
		period = fc.Int("p")
		port = fc.Int("ap")
	}

	tree := loadPostgreTree(connStr)

	createSaver(tree, nodeURL, privateKeyHex, contractAddress, period)

	createAndStartAPI(tree, port)
}
