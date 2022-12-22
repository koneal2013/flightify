package server

import (
	"github.com/gorilla/mux"
)

type Config struct {
	IsDevelopment   bool
	Port            int
	MiddlewareFuncs []mux.MiddlewareFunc
}
