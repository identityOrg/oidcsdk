package strategies

import (
	"log"
	"net/http"
	sdk "oauth2-oidc-sdk"
)

func DefaultLoggingErrorStrategy(err error, w http.ResponseWriter) {
	log.Print(err)
	w.WriteHeader(500)
	w.Header().Set(sdk.HeaderContentType, sdk.ContentTypeHtml)
	_, err = w.Write([]byte("unknown service error"))
	log.Print(err)
}
