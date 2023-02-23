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
	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))

	var err error
	db, err = ConnectDB(dbInfo)
	if err != nil {
		log.Fatal(err)
	}
}

func ConnectDB(dataSourceName string) (*gorp.DbMap, error) {
	postGresDb, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	if err = postGresDb.Ping(); err != nil {
		return nil, err
	}

	dbmap := &gorp.DbMap{Db: postGresDb, Dialect: gorp.PostgresDialect{}}
	return dbmap, nil
}

func GetDB() *gorp.DbMap {
	return db
}
