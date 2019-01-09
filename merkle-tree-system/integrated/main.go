package main

import (
	"./saver"
	"fmt"
	"github.com/LimeChain/merkletree"
	"github.com/LimeChain/merkletree/memory"
	"github.com/LimeChain/merkletree/postgres"
	merkleRestAPI "github.com/LimeChain/merkletree/restapi/baseapi"
	validateAPI "github.com/LimeChain/merkletree/restapi/validateapi"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"log"
	"net/http"
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
		treeRouter = validateAPI.MerkleTreeValidate(treeRouter, tree)
		r.Mount("/api/merkletree", treeRouter)
	})
	log.Fatal(http.ListenAndServe(":8080", router))
}

func createSaver(tree merkletree.MerkleTree) {
	treeSaver, err := saver.NewSaver("http://localhost:8545/", "7ab741b57e8d94dd7e1a29055646bafde7010f38a900f55bbd7647880faa6ee8", "0xa00f6A3a3D00979D7B7E23D7C7dF6CC7E255Ad88", tree)
	if err != nil {
		panic(err)
	}

	tx, err := treeSaver.TriggerSave()
	if err != nil {
		panic(err)
	}
	fmt.Println(tx)
}

func main() {
	connStr := "user=georgespasov dbname=postgres port=5432 sslmode=disable"
	tree := postgres.LoadMerkleTree(memory.NewMerkleTree(), connStr)

	createSaver(tree)
	createAndStartAPI(tree)

}
