package main

import (
	"log"

	"github.com/gustavohiroaki/personalfinance/infrastructure"
	"github.com/gustavohiroaki/personalfinance/routers"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	serverInstance := infrastructure.PrepareServer()
	routers := routers.InitRouter(serverInstance)
	infrastructure.InitServer(routers)
}
