package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/google/uuid"
)

// Transaction represents the financial event schema for $7T asset management scenarios [cite: 12]
type Transaction struct {
	ID        string  `json:"transaction_id"`
	Amount    float64 `json:"amount"`
	Currency  string  `json:"currency"`
	AccountID string  `json:"account_id"`
	Type      string  `json:"type"`
	Timestamp int64   `json:"timestamp"`
}

// Global pool to minimize GC overhead during "massive stream" simulation [cite: 6, 11]
var txPool = sync.Pool{
	New: func() interface{} {
		return &Transaction{}
	},
}

func generateLoad(workerID int, duration time.Duration) {
	ticker := time.NewTicker(10 * time.Microsecond) // Simulate high frequency
	stop := time.After(duration)

	for {
		select {
		case <-ticker.C:
			tx := txPool.Get().(*Transaction)
			tx.ID = uuid.New().String()
			tx.Amount = rand.Float64() * 50000
			tx.AccountID = fmt.Sprintf("ACT-%04d", rand.Intn(1000))
			tx.Type = []string{"DEBIT", "CREDIT", "TRANSFER"}[rand.Intn(3)]
			tx.Timestamp = time.Now().UnixMilli()

			// Simulate high-throughput processing
			_, _ = json.Marshal(tx)

			txPool.Put(tx)
		case <-stop:
			return
		}
	}
}

func main() {
	fmt.Println("🚀 FinPulse: Initializing Ultra-High-Throughput Generator...")
	var wg sync.WaitGroup

	// Demonstrate scaling across CPU cores [cite: 20]
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			generateLoad(id, 30*time.Second)
		}(i)
	}
	wg.Wait()
	fmt.Println("✅ Stream simulation complete.")
}
