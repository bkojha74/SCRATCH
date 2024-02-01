# SCRATCH
It means every thing written from scratch.

# About
This project is a scratchpad for me to write down my ideas and code snippets.
It is Backend project which is written in Golang.
It interacts with PostgreSQL database.

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