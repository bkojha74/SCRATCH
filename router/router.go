package router

import (
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

	s := mux.NewRouter().PathPrefix("/api").Subrouter()
	s.HandleFunc("/hello", helper.MessageHandler).Methods("GET")
	s.HandleFunc("/users", db.CreateUserHandler).Methods("POST")
	s.HandleFunc("/users", db.MiddlewareAuth(db.GetUserHandler)).Methods("GET")
	s.HandleFunc("/feeds", db.MiddlewareAuth(db.CreateFeedHandler)).Methods("POST")
	s.HandleFunc("/feeds", db.GetFeedHandler).Methods("GET")
	s.HandleFunc("/feed_follows", db.MiddlewareAuth(db.CreateFeedFollowHandler)).Methods("POST")
	s.HandleFunc("/feed_follows", db.MiddlewareAuth(db.GetFeedFollowsHandler)).Methods("GET")
	s.HandleFunc("/feed_follows/{feed_follow_id}", db.MiddlewareAuth(db.DeleteFeedFollowHandler)).Methods("DELETE")

	s.Use(middleware.HttpLogger)

	router.PathPrefix("/api").Handler(s)

	return router
}
