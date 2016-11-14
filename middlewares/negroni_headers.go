package middlewares

import (
	"net/http"

	"github.com/urfave/negroni"
)

func NewNegroniHeaders(key string, val string) negroni.Handler {
	return negroni.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
			w.Header().Add(key, val)
			next(w, r)
		},
	)
}
