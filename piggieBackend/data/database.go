package data

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

const (
	host   string = "localhost"
	port   int    = 5432
	dbname string = "piggie"
)

var (
	user     = os.Getenv("SQLUSER")
	password = os.Getenv("SQLPASSWORD")
)

func InitDB() {
	config := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	DB, err = sql.Open("postgres", config)
	if err != nil {
		log.Fatal("Error occured while opening connection to database server")
	}
}

func CloseDB() {
	err := DB.Close()
	if err != nil {
		log.Fatal("Error occured while closing connection to database server")
	}
}
