package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	port := os.Getenv("PORT")

	log.Printf("Application running on %s\n", port)
	log.Fatalln(http.ListenAndServe(":"+port, nil))
}
