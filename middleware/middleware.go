package middleware

/*import (
	"log"
	"net/http"
)

func HttpLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := log.New(log.Writer(), "", log.LstdFlags)
		logger.Println(r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}*/

import (
	"io"
	"log"
	"net/http"
)

// HttpLogger logs HTTP requests
func HttpLogger(writer io.Writer, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Use a custom logger with the provided writer
		logger := log.New(writer, "", 0)

		// Log the request details
		logger.Printf("%s %s\n", r.Method, r.URL.Path)

		// Pass the request to the next handler
		next.ServeHTTP(w, r)
	})
}
