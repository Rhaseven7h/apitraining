package main

import (
	"fmt"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/rhaseven7h/apitraining/controllers"
	"github.com/rhaseven7h/apitraining/middlewares"
	"github.com/urfave/negroni"
)

// interface http.Handler {
// 	ServeHTTP(w http.ResponseWriter, r *http.Request)
// }
//

// type MyHandler struct {
// 	myHTTPHandlerHandler MyHandlerFuncSignature
// }
//
// func (mh MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	mh.myHTTPHandlerHandler(w, r)
// }
//
// type MyHandlerFuncSignature func(http.ResponseWriter, *http.Request)
//
// func HandleFunc(mf MyHandlerFuncSignature) http.Handler {
// 	return MyHandler{mf}
// }

func main() {
	bindAddress := "0.0.0.0:9099"

	logger := logrus.New()

	m := mux.NewRouter()

	productsController := controllers.NewProductsController(45)
	m.HandleFunc("/products", productsController.List).Methods("GET")
	m.HandleFunc("/products/{id}", productsController.Get).Methods("GET")

	servicesController := controllers.NewServicesController("ooyala!", logger)
	m.HandleFunc("/services", servicesController.List)
	m.HandleFunc("/services/{id}", servicesController.Get)

	gh1 := middlewares.GlobalHeader("Content-Type", "application/json", m)
	gh2 := middlewares.GlobalHeader("Accept", "application/json", gh1)
	lm := middlewares.LoggingMiddleware(gh2)

	n := negroni.Classic()
	// n := negroni.New(
	// 	negroni.NewRecovery(),
	// 	negroni.NewLogger(),
	// 	negroni.NewStatic(http.Dir("./public")),
	// )
	n.Use(middlewares.NewNegroniHeaders("Newheader", "NewVal"))

	n.UseHandler(lm)

	fmt.Printf("Server started, listening at %s\n", bindAddress)
	http.ListenAndServe(bindAddress, n)
}
