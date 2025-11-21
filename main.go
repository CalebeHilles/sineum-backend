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

	defer func() {
		err := database.CloseDatabase()
		if err != nil {
			fmt.Println("Error closing the database:", err)
		}
	}()

	fmt.Println("Initializing server...")
	routes.HandleRequest()
}

func openDotEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Proceeding with production environment variables")
	}
}
