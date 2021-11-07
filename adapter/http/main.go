package http

import (
	"net/http"

	"github.com/DanielTrondoli/My-Finance-Server/adapter/http/actuator"
	"github.com/DanielTrondoli/My-Finance-Server/adapter/http/transaction"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Init blablabal
func Init() {
	http.HandleFunc("/transaction/create", transaction.CreateTransactions)
	http.HandleFunc("/transaction", transaction.GetTransaction)

	http.HandleFunc("/health", actuator.Health)

	http.HandleFunc("/metrics", promhttp.Handler().ServeHTTP)

	http.ListenAndServe(":8080", nil)
}
