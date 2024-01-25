package main

import (
	"log"
	"net/http"
	"os"

	"github.com/bkojha74/rssagg/router"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("$PORT must be set")
	}

	srv := &http.Server{
		Handler: router.NewRouter(),
		Addr:    ":" + portString,
	}

	log.Println("Server starting on PORT=", portString)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
