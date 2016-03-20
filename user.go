package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

type UserResource struct{}

func (u *UserResource) Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", u.listUsersAction).Methods(http.MethodGet)
	router.HandleFunc("/{userId}", u.getUserAction).Methods(http.MethodGet)

	return router
}

func (u *UserResource) PreferredPrefix() string {
	return "/users"
}

func (u *UserResource) listUsersAction(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	res.Header().Set("Content-Type", "application/json")
	res.Write([]byte("[]"))
}

func (u *UserResource) getUserAction(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusNotFound)
	res.Header().Set("Content-Type", "application/json")
	res.Write([]byte("{\"error\": \"This isn't implemented\"}"))
}
