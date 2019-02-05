package main

import (
	"encoding/json"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-chi/render"
)

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
