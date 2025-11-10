package main

import (
	"api-go/database"
	"api-go/routes"
	"fmt"
	"github.com/lpernett/godotenv"
	"log"
)

func main() {
	openDotEnv()

	database.ConnectToDatabase()
	fmt.Println("Initializing server...")
	routes.HandleRequest()
}

func openDotEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}
}
