package router

import (
	"net/http"
	"os"
	"time"

	"github.com/bkojha74/rssagg/backgroundjob"
	"github.com/bkojha74/rssagg/helper"
	"github.com/bkojha74/rssagg/middleware"
	"github.com/gorilla/mux"
)

func initDB() *helper.DBConfig {
	dbCfg := helper.DBConfig{}
	dbCfg.Init()
	return &dbCfg
}

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	db := initDB()

	go backgroundjob.StartScrapping(db.DB, 10, time.Minute)

	s := mux.NewRouter().PathPrefix("/api").Subrouter()
	s.HandleFunc("/hello", helper.MessageHandler).Methods("GET")
	s.HandleFunc("/users", db.CreateUserHandler).Methods("POST")
	s.HandleFunc("/users", db.MiddlewareAuth(db.GetUserHandler)).Methods("GET")
	s.HandleFunc("/feeds", db.MiddlewareAuth(db.CreateFeedHandler)).Methods("POST")
	s.HandleFunc("/feeds", db.GetFeedHandler).Methods("GET")
	s.HandleFunc("/feed_follows", db.MiddlewareAuth(db.CreateFeedFollowHandler)).Methods("POST")
	s.HandleFunc("/feed_follows", db.MiddlewareAuth(db.GetFeedFollowsHandler)).Methods("GET")
	s.HandleFunc("/feed_follows/{feed_follow_id}", db.MiddlewareAuth(db.DeleteFeedFollowHandler)).Methods("DELETE")
	s.HandleFunc("/posts", db.MiddlewareAuth(db.GetPostsForUserHandler)).Methods("GET")

	s.Use(func(next http.Handler) http.Handler {
		return middleware.HttpLogger(os.Stdout, next)
	})

	router.PathPrefix("/api").Handler(s)

	return router
}
