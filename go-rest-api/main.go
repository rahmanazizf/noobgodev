package main

import (
	"godev/go-rest-api/routers"
)

func main() {
	routers.StartServer().Run(":8382")
}
