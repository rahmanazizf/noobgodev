package database

import (
	"fmt"
	"godev/go-rest-api/helpers"
	"godev/go-rest-api/models"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBConnection() *gorm.DB {
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
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	helpers.CheckError(err)

	db.Exec("DROP TABLE IF EXISTS items") // items should be deleted first since there is foreign key that dependents to orders
	db.Exec("DROP TABLE IF EXISTS orders")
	log.Println("Dropped existing items and orders table")

	db.AutoMigrate(&models.Order{}, &models.Item{})

	return db

}
