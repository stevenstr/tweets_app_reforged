package main

import (
	"log"
	"net/http"

	"tweets.com/tweets/internal/controller/tweets"
	httphandler "tweets.com/tweets/internal/handler/http"
	"tweets.com/tweets/internal/repository/memory"
)

func main() {

	log.Println("Starting tweets microservice...")
	repo := memory.New()
	ctrl := tweets.New(repo)
	h := httphandler.New(ctrl)

	http.Handle("/tweets/list", http.HandlerFunc(h.HandleGetAllTweet))
	http.Handle("/tweets/get", http.HandlerFunc(h.HandleGetSingleTweet))
	http.Handle("/tweets/put", http.HandlerFunc(h.HandlePutSingleTweet))
	http.Handle("/tweets/time", http.HandlerFunc(h.HandleTime))

	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}

}
