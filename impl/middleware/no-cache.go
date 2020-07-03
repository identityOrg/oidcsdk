package middleware

import "net/http"

func NoCache(inner http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Cache-Control", "no-store")
		w.Header().Add("Cache-Control", "no-cache")
		inner(w, r)
	}
}
