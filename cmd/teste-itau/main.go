package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/OptionAndVoid/teste-itau/internal/api"
	"github.com/OptionAndVoid/teste-itau/internal/registry"
	"github.com/OptionAndVoid/teste-itau/pkg/logging"
	"github.com/OptionAndVoid/teste-itau/pkg/server"
	prom "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var pingCounter = prom.NewCounter(
	prom.CounterOpts{
		Name: "ping_request_count",
		Help: "No of request handled by Ping handler",
	},
)

func main() {

	// metrics setup
	prom.MustRegister(pingCounter)

	// logging setup
	logging.SetDefaultJSONLogger(os.Stdout, nil)
	slog.Info("Loggin setup de cria")

	is_tls := false
	if len(os.Args) == 0 {
		slog.Info("Running without tls")
	} else if len(os.Args) == 3 {
		is_tls = true
		slog.Info("Running with tls")
	} else {
		slog.Error("Please run with 0 args (for no tls) or with 2 args (for tls)")
		os.Exit(1)
	}

	// create "db" data structure
	registry := registry.NewTransactionRegistry()

	transactionsController := api.NewTransactionController(registry)
	statisticsController := api.NewStatisticsController(registry)

	// setup routes
	mux := http.NewServeMux()

	mux.HandleFunc("POST /transacao", transactionsController.PostTransaction)

	mux.HandleFunc("DELETE /transacao", transactionsController.DeleteTransaction)

	mux.HandleFunc("GET /estatistica", statisticsController.GetStatistics)

	mux.HandleFunc("GET /healtchcheck", func(w http.ResponseWriter, r *http.Request) {
		pingCounter.Inc()
		api.WriteJSON(w, http.StatusOK, map[string]string{"status": "up"})
		return
	})

	mux.Handle("/metrics", promhttp.Handler())

	// setup server
	server := server.NewServer().
		WithHost("localhost").
		WithPort(8080).
		WithMux(mux)

	// run api
	if is_tls {
		err := server.RunWithTls(os.Args[1], os.Args[2])
		if err != nil {
			slog.Error("tls run failed", "error", err.Error())
			os.Exit(1)
		}
	} else {
		err := server.Run()
		if err != nil {
			slog.Error("tls run failed", "error", err.Error())
			os.Exit(1)
		}
	}
}
