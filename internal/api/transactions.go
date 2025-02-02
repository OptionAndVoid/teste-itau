package api

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/OptionAndVoid/teste-itau/internal/registry"
)

type Transaction interface {
	PostTransaction(w http.ResponseWriter, r *http.Request)
	DeleteTransaction(w http.ResponseWriter, r *http.Request)
}

type TransactionController struct {
	registry *registry.TransactionRegistry
}

type TransactionDto struct {
	Value    *float64   `json:"valor"`
	DateTime *time.Time `json:"dataHora"`
}

func (trdto *TransactionDto) ToTransaction() (registry.Transaction, error) {
	if trdto == nil {
		return registry.Transaction{}, fmt.Errorf("nil pointer")
	}

	if trdto.Value == nil || *trdto.Value < 0 {
		return registry.Transaction{}, fmt.Errorf("value must be >= 0")
	}

	if trdto.DateTime == nil {
		return registry.Transaction{}, fmt.Errorf("date must not be nil")
	}

	if trdto.DateTime.After(time.Now()) {
		return registry.Transaction{}, fmt.Errorf("transaction must have happened in the past")
	}

	return registry.Transaction{
		Value:    *trdto.Value,
		DateTime: *trdto.DateTime,
	}, nil
}

func NewTransactionController(registry *registry.TransactionRegistry) *TransactionController {
	return &TransactionController{
		registry,
	}
}

func (tc *TransactionController) PostTransaction(w http.ResponseWriter, r *http.Request) {
	var body TransactionDto

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		slog.Error("bad request", "error", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if body.Value == nil || body.DateTime == nil {
		slog.Error("bad request", "error", "value or dataHora does not exist in body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	transaction, err := body.ToTransaction()
	if err != nil {
		slog.Error("conversion error", "error", err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	tc.registry.AddTransaction(transaction)
	w.WriteHeader(http.StatusCreated)
	return
}

func (tc *TransactionController) DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	tc.registry.Clear()
	w.WriteHeader(http.StatusOK)
	return
}
