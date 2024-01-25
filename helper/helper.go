package helper

import (
	"encoding/json"
	"log"
	"net/http"
)

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
