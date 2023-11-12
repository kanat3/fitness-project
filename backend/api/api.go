package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	var r *mux.Router
	r = mux.NewRouter()
	return r
}

func Status() http.Handler {
	var Status = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Router is working"))
	})
	return Status
}
