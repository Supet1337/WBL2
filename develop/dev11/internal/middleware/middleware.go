package middleware

import (
	"log"
	"net/http"
)

func Log(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("method", r.Method, "path", r.URL.EscapedPath())
		w.Header().Add("Content-Type", "application/json")
		next(w, r)
	}
}
