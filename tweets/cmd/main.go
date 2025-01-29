package main

import (
	"tweets.com/tweets/internal/controller/tweets"
	"tweets.com/tweets/internal/repository/memory"
)

func main() {

	repo := memory.New()
	ctrl := tweets.New(repo)

}
