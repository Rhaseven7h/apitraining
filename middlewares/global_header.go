package middlewares

import "net/http"

func GlobalHeader(key string, val string, oh http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add(key, val)
		oh.ServeHTTP(w, r)
	})
}
