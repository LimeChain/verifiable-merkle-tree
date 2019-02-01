package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"./saver"
	"github.com/LimeChain/merkletree"
	"github.com/LimeChain/merkletree/memory"
	"github.com/LimeChain/merkletree/postgres"
	merkleRestAPI "github.com/LimeChain/merkletree/restapi/baseapi"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
	"github.com/kelseyhightower/envconfig"
)

type Configuration struct {
	DBConnection    string `envconfig:"DB_CONNECTION" required:"true"`    // connection string for the postgres database
	EthNodeURL      string `envconfig:"ETH_NODE_URL" required:"true"`     // URL to the ethereum node to connect
	EthPrivateKey   string `envconfig:"ETH_PRIVATE_KEY" required:"true"`  // private key for the ethereum saver
	ContractAddress string `envconfig:"CONTRACT_ADDRESS" required:"true"` // address to the verifier contract
	Port            int    `default:"8080"`                               // port to run the API on
	Period          int    `default:"60"`                                 // period to try and save the new root
}

var config Configuration
var tokenAuth *jwtauth.JWTAuth

func init() {
	tokenAuth = jwtauth.New("HS256", []byte("secrettt"), nil)
}

func createAndStartAPI(tree merkletree.ExternalMerkleTree, port int) {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)

	router.Route("/v1", func(r chi.Router) {
		treeRouter := chi.NewRouter()

		treeRouter.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(tokenAuth))
			r.Use(jwtauth.Authenticator)

			r = merkleRestAPI.MerkleTreeInsert(r, tree)
		})

		treeRouter.Group(func(r chi.Router) {
			r.Post("/token", getToken())

			r = merkleRestAPI.MerkleTreeStatus(r, tree)
			r = merkleRestAPI.MerkleTreeHashes(r, tree)
		})

		r.Mount("/api/merkletree", treeRouter)
	})

	fmt.Printf("Starting REST Api at port %v\n", port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), router))
}

type tokenDataRequest struct {
	Data string `json:"data"`
}

type tokenDataResponse struct {
	MerkleAPIResponse
	Token string `json:"token,omitempty"`
}

type MerkleAPIResponse struct {
	Status bool   `json:"status"`
	Error  string `json:"error,omitempty"`
}

func getToken() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var b tokenDataRequest
		err := decoder.Decode(&b)

		if err != nil {
			render.JSON(w, r, tokenDataResponse{MerkleAPIResponse{false, err.Error()}, ""})
			return
		}

		if b.Data == "" {
			render.JSON(w, r, tokenDataResponse{MerkleAPIResponse{false, "Missing data field"}, ""})
			return
		}

		_, tokenString, _ := tokenAuth.Encode(jwt.MapClaims{"user_id": []byte(b.Data)})
		render.JSON(w, r, tokenDataResponse{MerkleAPIResponse{true, ""}, tokenString})
	}
}

func createSaver(tree merkletree.MerkleTree, nodeUrl, privateKeyHex, contractAddress string, periodSeconds int) {
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
			savedRoot, err := treeSaver.FetchRoot()
			if err != nil {
				fmt.Println("ERR: Could not save the tree root")
				fmt.Println(err.Error())
				time.Sleep(timeout)
				continue
			}

			if savedRoot == tree.Root() {
				fmt.Printf("Same root (%v) found in the chain. Skipping this iteration\n", savedRoot)
				time.Sleep(timeout)
				continue
			}

			treeLen := tree.Length()
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

func loadPostgreTree(connStr string) merkletree.FullMerkleTree {
	tree := postgres.LoadMerkleTree(memory.NewMerkleTree(), connStr)
	fmt.Printf("Merkle tree loaded. Length : %v, Root : %v\n", tree.Length(), tree.Root())
	return tree
}

func main() {
	envconfig.Process("configuration", &config)

	connStr := config.DBConnection
	tree := loadPostgreTree(connStr)

	nodeUrl := config.EthNodeURL
	privateKeyHex := config.EthPrivateKey
	contractAddress := config.ContractAddress
	period := config.Period
	createSaver(tree, nodeUrl, privateKeyHex, contractAddress, period)

	port := config.Port
	createAndStartAPI(tree, port)
}
