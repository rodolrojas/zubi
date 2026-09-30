package httpHandler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func NewHTTPHandler() (*mux.Router, error) {
	handler := mux.NewRouter()
	handler.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.Debugf("%+v", r)
		logrus.Debugf("%s", r.URL)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		_, _ = w.Write([]byte("Hello, world!"))
	})
	return handler, nil
}