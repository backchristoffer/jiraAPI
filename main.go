package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func loadenv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error! No .env file found: %s", err)
	}
}

func main() {
	loadenv()
	apikey := os.Getenv("JIRA_TOKEN")
	url := os.Getenv("JIRA_URL")
	fmt.Println("hello")
}
