package main

import (
	"log"
	"net/http"
	"os"

	"github.com/bkojha74/rssagg/router"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	/*feed, err := rss.UrlToFeed("https://timesofindia.indiatimes.com/rssfeedstopstories.cms")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(feed)*/

	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
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
