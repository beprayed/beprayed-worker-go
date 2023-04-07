package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-gorp/gorp"
	_ "github.com/lib/pq"
)

// type DB struct {
// 	*sql.DB
// }

var db *gorp.DbMap

func Init() {
	fmt.Println("Initializing...")

	dbInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"))

	fmt.Println(dbInfo)

	var err error
	db, err = ConnectDB(dbInfo)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("DB Connected")
}

func ConnectDB(dataSourceName string) (*gorp.DbMap, error) {
	fmt.Println("Connecting to DB...", dataSourceName)
	postGresDb, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		fmt.Println("Error connecting to DB...", err)
		return nil, err
	}

	fmt.Println("Pinging DB...")
	if err = postGresDb.Ping(); err != nil {
		fmt.Println("Error pinging DB...", err)
		return nil, err
	}

	fmt.Print("DB Connected?")
	dbmap := &gorp.DbMap{Db: postGresDb, Dialect: gorp.PostgresDialect{}}
	return dbmap, nil
}

func GetDB() *gorp.DbMap {
	return db
}
