package main

import (
	"log"

	"github.com/gustavohiroaki/personalfinance/infrastructure"
	"github.com/gustavohiroaki/personalfinance/routers"
	"github.com/joho/godotenv"
)

func main() {
	const Version = "0.1.0"
	err := godotenv.Load(".env")
	log.Println("Starting Personal Finance API version ", Version)
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	serverInstance := infrastructure.PrepareServer()
	routers := routers.InitRouter(serverInstance)
	infrastructure.InitServer(routers)
}
