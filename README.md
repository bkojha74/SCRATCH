# SCRATCH
It means every thing written from scratch.

# About
This project is a scratchpad for me to write down my ideas and code snippets.
It is Backend project which is written in Golang.
It interacts with PostgreSQL database.

# Software required
- Golang
- Keycloak (https://www.keycloak.org/)
- PostgreSQL with pgAdmin4 or any other GUI tool for database management.
- Sqlc (https://github.com/sqlc-dev/sqlc)
- Goose (https://github.com/pressly/goose)


# How to run
- Clone this repository
- Open terminal and go to the project directory
- Run `go mod tidy`
- Run `go mod vendor`
- Run `go run .`

# Structure

```
SCRATCH
├─ .gitignore
├─ go.mod
├─ go.sum
├─ helper
│  └─ helper.go
├─ main.go
├─ middleware
│  └─ middleware.go
├─ models
│  └─ models.go
├─ README.md
├─ router
│  └─ router.go
├─ sql
│  ├─ queries
│  │  └─ users.sql
│  └─ schema
│     └─ 001_users.sql
└─ sqlc.yaml

```