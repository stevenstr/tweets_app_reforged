package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/stevenstr/tweets_app_reforged/pkg/discovery"
	"github.com/stevenstr/tweets_app_reforged/pkg/discovery/consul"
	"github.com/stevenstr/tweets_app_reforged/tweets/internal/controller/tweets"
	httphandler "github.com/stevenstr/tweets_app_reforged/tweets/internal/handler/http"
	"github.com/stevenstr/tweets_app_reforged/tweets/internal/repository/memory"
)

const serviceName = "tweets"

func main() {
	var port int
	flag.IntVar(&port, "port", 8081, "api handler port")
	flag.Parse()

	log.Println("Starting tweets microservice...")
	registry, err := consul.NewRegistry("localhost:8500")
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	instanceID := discovery.GenerateInstanceID(serviceName)

	if err := registry.Register(ctx, instanceID, serviceName, fmt.Sprintf("localhost:%d", port)); err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			if err := registry.ReportHealthyState(
				instanceID, serviceName); err != nil {
				log.Println("Failed to report healthy state: " + err.Error())

			}
			time.Sleep(1 * time.Second)
		}
	}()
	defer registry.Deregister(ctx, instanceID, serviceName)

	repo := memory.New()
	ctrl := tweets.New(repo)
	h := httphandler.New(ctrl)
	mux := http.NewServeMux()

	mux.HandleFunc("/tweets/get", h.HandleGetSingleTweet)
	mux.HandleFunc("/tweets/post", h.HandlePostSingleTweet)
	mux.HandleFunc("/tweets/put", h.HandlePutSingleTweet)
	mux.HandleFunc("/tweets/delete", h.HandleDeleteSingleTweet)
	mux.HandleFunc("/tweets/list", h.HandleGetAllTweet)
	mux.HandleFunc("/tweets/time", h.HandleTime)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux); err != nil {
		log.Fatal(err)
	}
}
