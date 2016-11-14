package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/rhaseven7h/apitraining/loggingadapter"
)

type ServicesHandler struct {
	MyServiceID string
}

func NewServicesController(myServiceID string, logger loggingadapter.OOLogger) *ServicesHandler {
	logger.
		WithField("input_id", myServiceID).
		Info("creating services controller")
	return &ServicesHandler{
		MyServiceID: myServiceID,
	}
}

func (ph ServicesHandler) List(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "List of services!")
}

func (ph ServicesHandler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		logrus.
			WithField("id_received", vars["id"]).
			Error("error occurred with parameters")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, fmt.Sprintf("The ID requested was not an int: %s", vars["id"]))
		return
	}
	logrus.
		WithField("id_received", id).
		Info("error occurred with parameters")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, fmt.Sprintf("The ID requested was: %d", id))
	logrus.Info("Paso por aqui EN la accion")
}
