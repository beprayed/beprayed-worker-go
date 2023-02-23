package test

import (
	"log"

	"beprayed-worker-go/db"

	"github.com/joho/godotenv"
)

type TestUtil struct{}

func (t *TestUtil) InitDB() {
	dbInstance := db.GetDB()

	if dbInstance == nil {
		//Load the .env file
		err := godotenv.Load("../.env")
		if err != nil {
			log.Fatal("Error loading .env file, please create one in the root directory")
		}

		db.Init()
	}
}

func (t *TestUtil) ClearDataInPostgres() {
	db.GetDB().Exec("DELETE FROM public.pray_record")
}
