package backgroundjob

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/bkojha74/rssagg/internal/database"
	"github.com/bkojha74/rssagg/models/rss"
)

func StartScrapping(db *database.Queries, concurrency int, timeBetweenRequest time.Duration) {
	log.Printf("Scrapping on %v goroutines every %s duration", concurrency, timeBetweenRequest)
	ticker := time.NewTicker(timeBetweenRequest)
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(concurrency))
		if err != nil {
			fmt.Println("Error while getting next feeds to fetch:", err)
			continue
		}
		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)
			go scrapFeed(db, wg, feed)
		}
		wg.Wait()
	}

}

func scrapFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("Error while marking feed as fetched", err)
		return
	}

	rssFeed, err := rss.UrlToFeed(feed.Url)
	if err != nil {
		log.Println("Error while fetching feed:", err)
	}

	for _, item := range rssFeed.Channel.Item {
		log.Println("Found Post: ", item.Title, "on feed", feed.Name)
	}
	log.Printf("Feed %s collected. %v posts found", feed.Name, len(rssFeed.Channel.Item))
}
