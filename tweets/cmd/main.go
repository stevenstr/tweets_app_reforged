package main

import (
	"log"
	"net/http"

	"github.com/stevenstr/tweets_app_reforged/tweets/internal/controller/tweets"
	httphandler "github.com/stevenstr/tweets_app_reforged/tweets/internal/handler/http"
	"github.com/stevenstr/tweets_app_reforged/tweets/internal/repository/memory"
)

func main() {
	log.Println("Starting tweets microservice...")

	repo := memory.New()
	ctrl := tweets.New(repo)
	h := httphandler.New(ctrl)
	mux := http.NewServeMux()

	mux.HandleFunc("/tweets/list", h.HandleGetAllTweet)
	mux.HandleFunc("/tweets/get", h.HandleGetSingleTweet)
	mux.HandleFunc("/tweets/put", h.HandlePutSingleTweet)
	mux.HandleFunc("/tweets/time", h.HandleTime)

	if err := http.ListenAndServe(":8081", mux); err != nil {
		log.Fatal(err)
	}
}
