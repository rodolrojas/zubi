package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"rodolrojas.com/zubi/internal/common/structs"
	"rodolrojas.com/zubi/internal/handler"
)

type Router struct {
	RouteCollection RouteCollection
	Handler *mux.Router
}

func NewRouter(services *structs.BaseServerConfig) (*Router, error) {
	__handler, err := handler.NewHTTPHandler(); 
	if err != nil {
		logrus.Fatalf("Cannot initialize HTTP handler: %v", err)
	}
	var r = &Router{
		RouteCollection: *NewRouteCollection(services),
		Handler: __handler.GetRouter(),
	}
	r.SetupDefaultRoute()
	r.RouteCollection.setupRoutes(r.Handler)

	return r, nil
}

func (r *Router) GetHandler() *mux.Router {
	return r.Handler
}

func (r *Router) SetupDefaultRoute() {
	r.Handler.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("Welcome to Zubi API Gateway"))
	}).Methods("GET")
}