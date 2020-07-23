package middleware

import "net/http"

func NoCache(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Cache-Control", "no-store")
		w.Header().Add("Cache-Control", "no-cache")
		inner.ServeHTTP(w, r)
	})
}
