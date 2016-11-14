package middlewares

import (
	"net/http"

	"github.com/Sirupsen/logrus"
)

func LoggingMiddleware(oh http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.Info("Paso por aqui antes de la accion")
		oh.ServeHTTP(w, r)
		logrus.Info("Paso por aqui despues de la accion")
	})
}
