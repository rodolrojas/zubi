package handler

import (
	"github.com/gorilla/mux"
)

type HTTPHandler struct {
	Router *mux.Router
}
	
func NewHTTPHandler() (*HTTPHandler, error) {
	router := mux.NewRouter()
	return &HTTPHandler{
		Router: router,
	}, nil
}

func (h *HTTPHandler) GetRouter() *mux.Router {
	return h.Router
}