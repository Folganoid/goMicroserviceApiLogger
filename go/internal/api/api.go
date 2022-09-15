package api

import (
	"net/http"
)

func Start(config *Config) error {

	srv := newServer()

	return http.ListenAndServe(config.BindAddr, srv)
}
