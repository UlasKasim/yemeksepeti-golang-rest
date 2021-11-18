package server

import (
	stdlog "log"
	"net/http"
	"os"
	key_value "yemeksepeti-golang-rest/module/key_value"

	admissioncontrol "github.com/elithrar/admission-control"
	log "github.com/go-kit/kit/log"
)

func StartServer() {
	//Initialize key-value http server
	key_value.Initialize()

	var logger log.Logger
	// Logfmt is a structured, key=val logging format that is easy to read and parse
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	// Direct any attempts to use Go's log package to our structured logger
	stdlog.SetOutput(log.NewStdlibAdapter(logger))
	// Log the timestamp (in UTC) and the callsite (file + line number) of the logging
	// call for debugging in the future
	logger = log.With(logger, "Time", log.DefaultTimestampUTC, "loc", log.DefaultCaller)

	// Create an instance of our LoggingMiddleware with our configured logger
	loggingMiddleware := admissioncontrol.LoggingMiddleware(logger)
	loggedRouter := loggingMiddleware(http.DefaultServeMux)

	// Start our HTTP server
	err := http.ListenAndServe(":8080", loggedRouter)
	if err != nil {
		logger.Log("status", "fatal", "err", err)
		stdlog.Fatal(err)
	}
}
