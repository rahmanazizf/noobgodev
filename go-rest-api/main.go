package main

import (
	"godev/go-rest-api/controllers"
	"godev/go-rest-api/database"
	"godev/go-rest-api/routers"
)

func main() {
	conn := database.DBConnection()
	db := controllers.NewConnection(conn)
	controllers.EstablishConnection(db)
	routers.StartServer().Run(":8382")
}
