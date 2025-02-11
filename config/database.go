package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func InitDB() *sql.DB {
	// Get env variables loaded from main
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	//int data source name
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbUser, dbPass, dbHost, dbPort, dbName)

	//open connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	//PING to check connection
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}
