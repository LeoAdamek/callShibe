package main

import (
	"github.com/gorilla/mux"
)

type Resource interface {
	PreferredPrefix() string
	Router() *mux.Router
}
