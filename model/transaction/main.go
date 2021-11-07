package transaction

import "time"

type Transaction struct {
	Title     string    `jason:"title"`
	Amount    float32   `jason:"amount"`
	Type      int       `jason:"type"`
	CreatedAt time.Time `jason:"created_at"`
}

type Transactions []Transaction
