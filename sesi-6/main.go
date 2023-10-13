package main

import (
	"database/sql"
	"fmt"
	"godev/sesi-6/cmd"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Fisikaasyik123"
	dbname   = "godev"
)

func main() {
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlConn)
	cmd.CheckError(err)
	defer db.Close()
	err = db.Ping()
	cmd.CheckError(err)
	log.Println("Connected to the database!")

	// creating table
	cmd.CreateTable(db)
}
