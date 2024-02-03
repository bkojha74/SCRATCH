package rss

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUrlToFeed(t *testing.T) {
	// Mocking a simple RSS XML response
	mockXML := `
	<rss version="2.0">
		<channel>
			<title>Test RSS Feed</title>
			<link>https://example.com</link>
			<description>Testing RSS Feed</description>
			<language>en-us</language>
			<item>
				<title>Item 1</title>
				<link>https://example.com/item1</link>
				<description>Description for Item 1</description>
				<pubDate>Mon, 02 Jan 2006 15:04:05 MST</pubDate>
			</item>
			<item>
				<title>Item 2</title>
				<link>https://example.com/item2</link>
				<description>Description for Item 2</description>
				<pubDate>Tue, 03 Jan 2006 15:04:05 MST</pubDate>
			</item>
		</channel>
	</rss>`

	// Setup a mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/xml")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockXML))
	}))
	defer server.Close()

	// Call the function with the mock server's URL
	rssFeed, err := UrlToFeed(server.URL)
	assert.NoError(t, err, "Error should be nil")

	// Assertions on the RSS feed structure
	assert.Equal(t, "Test RSS Feed", rssFeed.Channel.Title)
	assert.Equal(t, "https://example.com", rssFeed.Channel.Link)
	assert.Equal(t, "Testing RSS Feed", rssFeed.Channel.Description)
	assert.Equal(t, "en-us", rssFeed.Channel.Language)

	// Assertions on the first item
	assert.Len(t, rssFeed.Channel.Item, 2) // Assuming two items in the mock data
	assert.Equal(t, "Item 1", rssFeed.Channel.Item[0].Title)
	assert.Equal(t, "https://example.com/item1", rssFeed.Channel.Item[0].Link)
	assert.Equal(t, "Description for Item 1", rssFeed.Channel.Item[0].Description)
	// Add more assertions based on your expected data and structure

	// Assertions on the second item
	assert.Equal(t, "Item 2", rssFeed.Channel.Item[1].Title)
	assert.Equal(t, "https://example.com/item2", rssFeed.Channel.Item[1].Link)
	assert.Equal(t, "Description for Item 2", rssFeed.Channel.Item[1].Description)
	// Add more assertions based on your expected data and structure
}
