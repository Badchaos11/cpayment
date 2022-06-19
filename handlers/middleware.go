package handlers

import (
	"log"
	"net/http"
)

func (t *Transactions) MiddlewareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h := r.Header["Authorization"]
		t.l.Println(h)
		if h[0] != token {
			log.Println("Invalid authorization header received")
			return
		}
		t.l.Println("Token allowed, go to next handler")
		next.ServeHTTP(w, r)
	})
}
