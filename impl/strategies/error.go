package strategies

import (
	"log"
	"net/http"
)

func DefaultLoggingErrorStrategy(err error, w http.ResponseWriter) {
	log.Print(err)
	w.WriteHeader(500)
	w.Header().Set("content-type", "text/html")
	_, err = w.Write([]byte("unknown service error"))
	log.Print(err)
}
