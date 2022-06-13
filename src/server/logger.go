package server

import (
	"log"
	"net/http"
)

func (this *Server) Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s][HANDLER] New request %s '%s'.", this.Protocol, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}
