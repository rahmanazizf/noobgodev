package main

import (
	"database/sql"
	"fmt"
	"godev/sesi-6/cmd"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	var (
		host     = os.Getenv("HOST")
		port, _  = strconv.Atoi(os.Getenv("PORT"))
		user     = os.Getenv("USER")
		password = os.Getenv("PASSWORD")
		dbname   = os.Getenv("DBNAME")
	)
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlConn)
	cmd.CheckError(err)
	defer db.Close()
	err = db.Ping()
	cmd.CheckError(err)
	log.Println("Connected to the database!")

	// creating table
	cmd.CreateTable(db)

	// create product
	cmd.CreateProduct(db, "Indomie")
}
