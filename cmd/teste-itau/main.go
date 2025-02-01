package main

import (
	"log/slog"
	"os"

	"github.com/OptionAndVoid/teste-itau/internal/server"
	"github.com/OptionAndVoid/teste-itau/pkg/logging"
)

func main() {
	// logging setup
	logging.SetDefaultJSONLogger(os.Stdout, nil)
	slog.Info("Loggin setup de cria")

	// create "db" data structure

	// setup routes
	server := server.NewServer().WithHost("localhost").WithPort(8080)

	// run api
	server.Run()
}
