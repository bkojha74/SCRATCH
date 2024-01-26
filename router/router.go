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
	s.Use(middleware.HttpLogger)

	router.PathPrefix("/api").Handler(s)

	return router
}
