package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	fmt.Println("Hello World")
}
