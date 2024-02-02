# SCRATCH
As name suggest, it means every thing written from scratch.

# History
In the dynamic landscape of the internet, RSS feeds have emerged as an essential tool for content consumption and information dissemination. Born out of the need for a more efficient way to keep up with the ever-expanding online content, Really Simple Syndication (RSS) has revolutionized the way we access and digest information.

Initially introduced in the late 1990s, RSS feeds quickly gained popularity as a standardized format for delivering updated content from various sources. Websites, blogs, and news outlets embraced this technology, allowing users to aggregate headlines, summaries, and links to full articles in a single, user-friendly interface.

RSS feeds played a pivotal role in empowering users to personalize their online experience, enabling them to stay informed without navigating multiple websites. While their prominence may have waned in the era of social media, RSS feeds continue to provide a focused and efficient means for information enthusiasts to curate and manage their digital content.

Today, RSS feeds persist as a testament to the enduring relevance of simplicity in the rapidly evolving digital realm. Despite shifts in online consumption habits, the fundamental principles of RSS – simplicity, efficiency, and accessibility – have left an indelible mark on the way we engage with the vast ocean of information available on the internet.

As a developer, it's important to stay abreast of the latest trends and best practices in web development. One such trend is the rise of serverless architectures, which have gained significant traction in recent years. Serverless architectures allow developers to focus on building and scaling applications without the need for complex server management.

# About
This project is a scratchpad for me to write down my ideas and code snippets.
It is Backend project which is written in Golang.
It aggregate data from RSS Feeds using REST APIs.
It uses Keycloak for authentication and authorization.
It uses PostgreSQL for database.
It uses Sqlc for database migration.    
It uses Goose for database migration.

# Software required
- Golang
- Keycloak (https://www.keycloak.org/)
- PostgreSQL with pgAdmin4 or any other GUI tool for database management (https://www.postgresql.org/)
- Sqlc (https://github.com/sqlc-dev/sqlc)
- Goose (https://github.com/pressly/goose)


# How to run
- Clone this repository
- Open terminal and go to the project directory
- go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
- go install github.com/pressly/goose/v3/cmd/goose@latest
- go get github.com/lib/pq
- go mod tidy
- go mod vendor
- cd sql/schema
- goose postgres <db_url> up
- cd ../..
- sqlc generate  #it will create folder internal/database and some go files in it
- go build
- ./SCRATCH

# Structure

```
SCRATCH
├─ .gitignore
├─ auth
│  └─ auth.go
├─ backgroundjob
│  └─ scrapper.go
├─ go.mod
├─ go.sum
├─ helper
│  └─ helper.go
├─ internal
│  └─ database
│     ├─ db.go
│     ├─ feeds.sql.go
│     ├─ feed_follows.sql.go
│     ├─ models.go
│     ├─ posts.sql.go
│     └─ users.sql.go
├─ main.go
├─ middleware
│  └─ middleware.go
├─ models
│  ├─ models.go
│  └─ rss
│     └─ rss.go
├─ README.md
├─ router
│  └─ router.go
├─ sql
│  ├─ queries
│  │  ├─ feeds.sql
│  │  ├─ feed_follows.sql
│  │  ├─ posts.sql
│  │  └─ users.sql
│  └─ schema
│     ├─ 001_users.sql
│     ├─ 002_users_apikey.sql
│     ├─ 003_feeds.sql
│     ├─ 004_feed_follows.sql
│     ├─ 005_feeds_lastfetchedat.sql
│     └─ 006_posts.sql
└─ sqlc.yaml

```