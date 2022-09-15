package api

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"net/http"
)

type server struct {
	router *mux.Router
}

func newServer() *server {
	s := &server {
		router: mux.NewRouter(),
	}

	s.configureRouter()
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {

	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
	s.router.HandleFunc("/test", s.handleTest()).Methods("GET")

}
