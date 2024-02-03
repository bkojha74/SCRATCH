package models

import (
	"database/sql"
	"testing"
	"time"

	"github.com/bkojha74/rssagg/internal/database"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestDatabaseUserMap(t *testing.T) {
	dbUser := database.User{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      "Bipin",
		ApiKey:    "some-api-key",
	}

	mappedUser := DatabaseUserMap(dbUser)

	assert.Equal(t, dbUser.ID, mappedUser.ID)
	assert.Equal(t, dbUser.CreatedAt, mappedUser.CreatedAt)
	assert.Equal(t, dbUser.UpdatedAt, mappedUser.UpdatedAt)
	assert.Equal(t, dbUser.Name, mappedUser.Name)
	assert.Equal(t, dbUser.ApiKey, mappedUser.APIKey)
}

func TestDatabaseFeederMap(t *testing.T) {
	dbFeed := database.Feed{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      "Kumar",
		Url:       "https://kumar.com",
		UserID:    uuid.New(),
	}

	mappedFeed := DatabaseFeederMap(dbFeed)

	assert.Equal(t, dbFeed.ID, mappedFeed.ID)
	assert.Equal(t, dbFeed.CreatedAt, mappedFeed.CreatedAt)
	assert.Equal(t, dbFeed.UpdatedAt, mappedFeed.UpdatedAt)
	assert.Equal(t, dbFeed.Name, mappedFeed.Name)
	assert.Equal(t, dbFeed.Url, mappedFeed.Url)
	assert.Equal(t, dbFeed.UserID, mappedFeed.UserID)
}

func TestDatabaseFeedersMap(t *testing.T) {
	dbFeeds := []database.Feed{
		{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      "Ojha",
			Url:       "https://ojha.com",
			UserID:    uuid.New(),
		}, {
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      "Bko",
			Url:       "https://bko.com",
			UserID:    uuid.New(),
		},
	}

	mappedFeed := DatabaseFeedersMap(dbFeeds)
	assert.Len(t, mappedFeed, len(dbFeeds))
	for i, dbFeed := range dbFeeds {
		assert.Equal(t, dbFeed.ID, mappedFeed[i].ID)
		assert.Equal(t, dbFeed.CreatedAt, mappedFeed[i].CreatedAt)
		assert.Equal(t, dbFeed.UpdatedAt, mappedFeed[i].UpdatedAt)
		assert.Equal(t, dbFeed.Name, mappedFeed[i].Name)
		assert.Equal(t, dbFeed.Url, mappedFeed[i].Url)
	}
}

func TestDatabaseFeedFollowMap(t *testing.T) {
	dbFeedFollow := database.FeedFollow{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    uuid.New(),
		FeedID:    uuid.New(),
	}

	mappedFeedFollow := DatabaseFeedFollowMap(dbFeedFollow)

	assert.Equal(t, dbFeedFollow.ID, mappedFeedFollow.ID)
	assert.Equal(t, dbFeedFollow.CreatedAt, mappedFeedFollow.CreatedAt)
	assert.Equal(t, dbFeedFollow.UpdatedAt, mappedFeedFollow.UpdatedAt)
	assert.Equal(t, dbFeedFollow.FeedID, mappedFeedFollow.FeedID)
	assert.Equal(t, dbFeedFollow.UserID, mappedFeedFollow.UserID)
}

func TestDatabaseFeedFollowersMap(t *testing.T) {
	dbFeedFollows := []database.FeedFollow{
		{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserID:    uuid.New(),
			FeedID:    uuid.New(),
		}, {
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserID:    uuid.New(),
			FeedID:    uuid.New(),
		},
	}

	mappedFeedFollows := DatabaseFeedFollowersMap(dbFeedFollows)
	assert.Len(t, mappedFeedFollows, len(dbFeedFollows))
	for i, dbFeedFollow := range dbFeedFollows {
		assert.Equal(t, dbFeedFollow.ID, mappedFeedFollows[i].ID)
		assert.Equal(t, dbFeedFollow.CreatedAt, mappedFeedFollows[i].CreatedAt)
		assert.Equal(t, dbFeedFollow.UpdatedAt, mappedFeedFollows[i].UpdatedAt)
		assert.Equal(t, dbFeedFollow.FeedID, mappedFeedFollows[i].FeedID)
		assert.Equal(t, dbFeedFollow.UserID, mappedFeedFollows[i].UserID)
	}
}

func TestDatabasePostMap(t *testing.T) {
	dbPost := database.Post{
		ID:          uuid.New(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Title:       "Hello World",
		Description: sql.NullString{String: "This is a test post", Valid: true},
		PublishedAt: time.Now(),
		Url:         "https://kumar.com",
		FeedID:      uuid.New(),
	}

	mappedPost := DatabasePostMap(dbPost)

	assert.Equal(t, dbPost.ID, mappedPost.ID)
	assert.Equal(t, dbPost.CreatedAt, mappedPost.CreatedAt)
	assert.Equal(t, dbPost.UpdatedAt, mappedPost.UpdatedAt)
	assert.Equal(t, dbPost.Title, mappedPost.Title)
	assert.EqualValues(t, dbPost.Description.String, *mappedPost.Description)
	assert.Equal(t, dbPost.PublishedAt, mappedPost.PublishedAt)
	assert.Equal(t, dbPost.Url, mappedPost.Url)
	assert.Equal(t, dbPost.FeedID, mappedPost.FeedID)
}

func TestDatabasePostsMap(t *testing.T) {
	dbPostsMap := []database.Post{
		{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       "Hello",
			Description: sql.NullString{String: "This is a test hello post", Valid: true},
			PublishedAt: time.Now(),
			Url:         "https://hello.com",
			FeedID:      uuid.New(),
		}, {
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       "World",
			Description: sql.NullString{String: "This is a test world post", Valid: true},
			PublishedAt: time.Now(),
			Url:         "https://world.com",
			FeedID:      uuid.New(),
		},
	}

	mappedPost := DatabasePostsMap(dbPostsMap)
	assert.Len(t, mappedPost, len(dbPostsMap))
	for i, dbPost := range dbPostsMap {
		assert.Equal(t, dbPost.ID, mappedPost[i].ID)
		assert.Equal(t, dbPost.CreatedAt, mappedPost[i].CreatedAt)
		assert.Equal(t, dbPost.UpdatedAt, mappedPost[i].UpdatedAt)
		assert.Equal(t, dbPost.Title, mappedPost[i].Title)
		assert.EqualValues(t, dbPost.Description.String, *mappedPost[i].Description)
		assert.Equal(t, dbPost.PublishedAt, mappedPost[i].PublishedAt)
		assert.Equal(t, dbPost.Url, mappedPost[i].Url)
		assert.Equal(t, dbPost.FeedID, mappedPost[i].FeedID)
	}
}
