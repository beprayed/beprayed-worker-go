package main

import (
	"beprayed-worker-go/db"
	"beprayed-worker-go/worker"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// load env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error: failed to load the env file")
	}

	db.Init()
	worker.Init()
}
