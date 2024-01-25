package router

import (
	"github.com/bkojha74/rssagg/helper"
	"github.com/bkojha74/rssagg/middleware"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	s := mux.NewRouter().PathPrefix("/api").Subrouter()
	s.HandleFunc("/hello", helper.MessageHandler).Methods("GET")
	s.Use(middleware.HttpLogger)

	router.PathPrefix("/api").Handler(s)

	return router
}
