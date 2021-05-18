package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("/home/adam/go/workspace/src/github.com/adamszpilewicz/bookstore_users-api/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// USER := os.Getenv("user")
	// PASSWORD := os.Getenv("password")
	// DBNAME := os.Getenv("dbname")
	HOST := os.Getenv("host")
	// PORT := os.Getenv("port")
	fmt.Println(HOST)
}
