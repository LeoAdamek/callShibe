package main

import "flag"

type AppParams struct {
	Verbose bool
	Bind    string
	Logfile string
}

func GetOptions() *AppParams {

	verbosePtr := flag.Bool("v", false, "Enable Verbose Output")
	bindPtr := flag.String("b", ":8008", "Bind Address")
	logfilePtr := flag.String("l", "STDOUT", "Logfile")

	flag.Parse()

	return &AppParams{
		Verbose: *verbosePtr,
		Bind:    *bindPtr,
		Logfile: *logfilePtr,
	}
}
