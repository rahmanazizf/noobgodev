package connection

import (
	"fmt"
	"godev/go-rest-api/models"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// TODO: migrasi gorm model ke database?
// // TODO: driver postgresql dalam gorm
// func DBConnection() *sql.DB {
// 	if err := godotenv.Load(); err != nil {
// 		log.Fatalf("Error loading .env file: %v", err)
// 	}
// 	var (
// 		host     = os.Getenv("HOST")
// 		port, _  = strconv.Atoi(os.Getenv("PORT"))
// 		user     = os.Getenv("USER")
// 		password = os.Getenv("PASSWORD")
// 		dbname   = os.Getenv("DBNAME")
// 	)
// 	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
// 	db, err := sql.Open("postgres", psqlConn)
// 	CheckError(err)
// 	defer db.Close()
// 	err = db.Ping()
// 	CheckError(err)
// 	// log.Println("Connected to the database!")
// 	return db
// }

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

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
	CheckError(err)

	defer db.AutoMigrate(&models.Order{}, &models.Item{})

	return db

}
