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
	server.Run()
}
