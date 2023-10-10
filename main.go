package main

import (
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	server := NewAPIServer(nil, ":3000")
	server.Run()
}
