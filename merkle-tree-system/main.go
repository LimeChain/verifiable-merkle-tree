package main

import (
	"./saver"
	"fmt"
	"github.com/LimeChain/merkletree"
	"github.com/LimeChain/merkletree/memory"
	"github.com/LimeChain/merkletree/postgres"
	merkleRestAPI "github.com/LimeChain/merkletree/restapi/baseapi"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"log"
	"net/http"
	"time"
)

func createAndStartAPI(tree merkletree.ExternalMerkleTree) {
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
		treeRouter = merkleRestAPI.MerkleTreeStatus(treeRouter, tree)
		treeRouter = merkleRestAPI.MerkleTreeInsert(treeRouter, tree)
		treeRouter = merkleRestAPI.MerkleTreeHashes(treeRouter, tree)
		r.Mount("/api/merkletree", treeRouter)
	})
	log.Fatal(http.ListenAndServe(":8080", router))
}

func createSaver(tree merkletree.MerkleTree) {
	treeSaver, err := saver.NewSaver(
		"http://localhost:8545/",
		"7ab741b57e8d94dd7e1a29055646bafde7010f38a900f55bbd7647880faa6ee8",
		"0x9eD274314f0fB37837346C425D3cF28d89ca9599",
		tree)
	if err != nil {
		panic(err)
	}

	go func() {
		len := 0
		timeout := 15 * time.Second
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
					fmt.Printf("Mined transaction hash (%v)\n", tx)
					len = treeLen
				}
			} else {
				fmt.Println("No changes to submit. Skipping this iteration")
			}
			time.Sleep(timeout)
		}
	}()

}

func main() {
	connStr := "user=georgespasov dbname=postgres port=5432 sslmode=disable"
	tree := postgres.LoadMerkleTree(memory.NewMerkleTree(), connStr)

	createSaver(tree)
	createAndStartAPI(tree)
	fmt.Println("Rest API Started")
}
