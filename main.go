package main

import (
	"beprayed-worker-go/db"
	"beprayed-worker-go/worker"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Initializing the worker..1")
	fmt.Println(os.Getenv("DB_NAME"))
	fmt.Println("Initializing the worker..2")
	fmt.Println(os.Getenv("DB_HOST"))
	fmt.Println("Initializing the worker..3")

	// load env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error: failed to load the env file")
	}

	fmt.Println("Initializing the worker..4")
	fmt.Println(os.Getenv("DB_NAME"))
	fmt.Println("Initializing the worker..5")
	fmt.Println(os.Getenv("DB_HOST"))
	fmt.Println("Initializing the worker..6")

	db.Init()
	fmt.Println("Initializing the worker..7")
	worker.Init()
	fmt.Println("Initialized")
}
