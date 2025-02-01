package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/OptionAndVoid/teste-itau/internal/api"
	"github.com/OptionAndVoid/teste-itau/internal/server"
	"github.com/OptionAndVoid/teste-itau/pkg/logging"
)

func main() {
	// logging setup
	logging.SetDefaultJSONLogger(os.Stdout, nil)
	slog.Info("Loggin setup de cria")

	// create "db" data structure

	transacaoController := api.TransacaoController{}

	// setup routes
	mux := http.NewServeMux()

	mux.HandleFunc("GET /transacao", transacaoController.GetTransacao)

	mux.HandleFunc("POST /transacao", transacaoController.PostTransacao)

	mux.HandleFunc("DELETE /transacao", transacaoController.DeleteTransacao)

	// setup server
	server := server.NewServer().
		WithHost("localhost").
		WithPort(8080).
		WithMux(mux)

	// run api
	server.Run()
}
