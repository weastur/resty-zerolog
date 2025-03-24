// main package is used to demonstrate how to use resty-zerolog with resty
package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	restyzerolog "github.com/weastur/resty-zerolog"
	"resty.dev/v3"
)

const ComponentCtxKey = "component"

func main() {
	// Configure zerolog
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.With().Str(ComponentCtxKey, "core").Logger()

	log.Info().Msg("Starting")

	client := resty.New()
	defer client.Close()

	// Configure separate logger for resty. Not necessary, but can be useful
	restyLogger := log.With().Str(ComponentCtxKey, "resty").Logger()

	client.EnableDebug()
	client.EnableTrace()
	// Now resty will use zerolog logger
	client.SetLogger(restyzerolog.New(restyLogger))

	_, _ = client.R().Get("https://httpbin.org/get")

	log.Info().Msg("Finished")
}
