package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/nathanaelcunningham/gothBase/internal/env"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	port := env.GetInt("PORT", 3000)

	routes := Routes()

	log.Printf("Starting server on port %d", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), routes); err != nil {
		log.Fatal(err)
	}
}
