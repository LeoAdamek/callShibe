package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)

var Params *AppParams

func main() {
	Params = GetOptions()

	server := NewServer(Params)

	err := server.ServeForever()

	if err != nil {
		log.WithError(err).Fatalln("Catastrophic Failure!")
		os.Exit(130)
	}

}
