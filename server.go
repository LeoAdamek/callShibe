package main

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"net/http"
)

type Server struct {
	Id      uint32
	logger  *log.Logger
	bind    string
	http    *http.Server
	handler http.Handler
}

func NewServer(params *AppParams) *Server {
	logger := log.New()

	if params.Verbose {
		logger.Level = log.DebugLevel
	}

	server := &Server{
		logger: logger,
		bind:   params.Bind,
		Id:     rand.Uint32(),
	}

	server.Initialize()

	return server
}

func (s *Server) Initialize() {
	router := mux.NewRouter()

	userResource := UserResource{}
	router.Handle(userResource.PreferredPrefix(), userResource.Router())

	s.handler = router

	s.http = &http.Server{
		Addr:    s.bind,
		Handler: s.handler,
	}
}

func (s *Server) ServeForever() error {
	s.logger.WithFields(log.Fields{
		"server_id": s.Id,
		"bind":      s.bind,
	}).Infoln("Starting Server")

	return s.http.ListenAndServe()
}
