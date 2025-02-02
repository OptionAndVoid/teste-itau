package api

import (
	"math"
	"net/http"
	"time"

	"github.com/OptionAndVoid/teste-itau/internal/registry"
)

type Statistics interface {
	GetStatistics(w http.ResponseWriter, r *http.Request)
}

type RegistryStats struct {
	Count int     `json:"count"`
	Sum   float64 `json:"sum"`
	Avg   float64 `json:"avg"`
	Min   float64 `json:"min"`
	Max   float64 `json:"max"`
}

type StatisticsController struct {
	registry *registry.TransactionRegistry
}

func NewStatisticsController(registry *registry.TransactionRegistry) *StatisticsController {
	return &StatisticsController{
		registry,
	}
}

func (sc *StatisticsController) GetStatistics(w http.ResponseWriter, r *http.Request) {
	transactions := sc.registry.GetTransactionsInInterval(time.Minute)

	count := len(transactions)
	if count == 0 {
		WriteJSON(w, http.StatusOK, RegistryStats{})
		return
	}

	var sum float64
	var min float64 = math.MaxFloat64
	var max float64 = 0
	var avg float64

	for _, it := range transactions {
		sum += it.Value
		if it.Value < min {
			min = it.Value
		}
		if it.Value > max {
			max = it.Value
		}
	}
	if sum == 0 {
		avg = 0
	} else {
		avg = sum / float64(count)
	}

	registryStats := RegistryStats{
		Count: count,
		Sum:   sum,
		Avg:   avg,
		Min:   min,
		Max:   max,
	}

	WriteJSON(w, http.StatusOK, registryStats)
	return
}
