package strategies

import (
	"log"
	"net/http"
	sdk "oidcsdk"
)

func DefaultLoggingErrorStrategy(err error, w http.ResponseWriter) {
	log.Print(err)
	w.Header().Set(sdk.HeaderContentType, sdk.ContentTypeHtml)
	w.WriteHeader(500)
	_, err = w.Write([]byte("unknown service error"))
	log.Print(err)
}
