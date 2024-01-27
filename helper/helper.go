package helper

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/bkojha74/rssagg/internal/auth"
	"github.com/bkojha74/rssagg/internal/database"
	"github.com/bkojha74/rssagg/models"
	"github.com/google/uuid"
)

type DBConfig struct {
	DB *database.Queries
}

func (dbCfg *DBConfig) Init() {
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found in the environment")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Can't connect to database", err)
	}

	dbCfg.DB = database.New(conn)
}

func MessageHandler(w http.ResponseWriter, r *http.Request) {
	RespondWithJson(w, http.StatusOK, "Hello, World!")
}

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Internal Server Error", msg)
	}

	response, _ := json.Marshal(struct {
		Code    int    `json:"code"`
		Message string `json:"error"`
	}{
		Code:    code,
		Message: msg,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func RespondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (apiCfg *DBConfig) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON:%s", err.Error()))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Could not ceate user:%s", err.Error()))
		return
	}

	RespondWithJson(w, http.StatusCreated, models.DatabaseUserMap(user))
}

func (apiCfg *DBConfig) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetApiKey(r.Header)
	if err != nil {
		RespondWithError(w, http.StatusForbidden, fmt.Sprintf("Auth error:%s", err.Error()))
		return
	}

	user, err := apiCfg.DB.GetUser(r.Context(), apiKey)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Could not get user:%s", err.Error()))
		return
	}

	RespondWithJson(w, http.StatusOK, models.DatabaseUserMap(user))
}
