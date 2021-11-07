package http

import (
	"net/http"

	"github.com/DanielTrondoli/My-Finance-Server/adapter/http/transaction"
)

func Init() {
	http.HandleFunc("/transaction/create", transaction.CreateTransactions)
	http.HandleFunc("/transaction", transaction.GetTransaction)

	http.ListenAndServe(":8080", nil)
}
