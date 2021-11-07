package transaction

import "time"

// Transaction blablabla
type Transaction struct {
	Title     string    `jason:"title"`
	Amount    float32   `jason:"amount"`
	Type      int       `jason:"type"`
	CreatedAt time.Time `jason:"created_at"`
}

// Transactions blablabla
type Transactions []Transaction
