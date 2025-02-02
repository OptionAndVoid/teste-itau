package registry

import (
	"sync"
	"time"
)

type Transaction struct {
	Value    float64
	DateTime time.Time
}

type TransactionRegistry struct {
	transactionBuffer []Transaction
	mutex             sync.Mutex
}

func NewTransactionRegistry() *TransactionRegistry {
	return &TransactionRegistry{
		transactionBuffer: make([]Transaction, 0),
	}
}

func (tr *TransactionRegistry) AddTransaction(newTransaction Transaction) error {
	tr.mutex.Lock()
	defer tr.mutex.Unlock()
	tr.transactionBuffer = append(tr.transactionBuffer, newTransaction)
	return nil
}

func (tr *TransactionRegistry) Clear() error {
	tr.mutex.Lock()
	defer tr.mutex.Unlock()
	tr.transactionBuffer = make([]Transaction, 0)
	return nil
}

func (tr *TransactionRegistry) GetTransactionsInInterval(overTheLast time.Duration) []Transaction {
	tr.mutex.Lock()
	defer tr.mutex.Unlock()
	now := time.Now()

	transactionsInInterval := make([]Transaction, 0)
	for _, it := range tr.transactionBuffer {
		if now.Sub(it.DateTime) < overTheLast {
			transactionsInInterval = append(transactionsInInterval, it)
		}
	}
	return transactionsInInterval
}
