package main

import (
	"godev/go-rest-api/connection"
	"godev/go-rest-api/controllers"
	"godev/go-rest-api/routers"
)

func main() {
	conn := connection.DBConnection()
	db := controllers.NewConnection(conn)
	controllers.EstablishConnection(db)
	routers.StartServer().Run(":8382")
}
