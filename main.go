package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/DanielTrondoli/My-Finance-Server/transaction"
)

func main() {

	http.HandleFunc("/transaction/create", createTransactions)
	http.HandleFunc("/transaction", getTransaction)

	http.ListenAndServe(":8080", nil)

}

func getTransaction(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-type", "application/json")

	var layout = "2006-01-02T15:04:05"

	dateRecived, _ := time.Parse(layout, "1992-11-07T16:52:00")

	var transactions = transaction.Transactions{
		transaction.Transaction{
			Title:     "Salario",
			Amount:    4400,
			Type:      0,
			CreatedAt: dateRecived,
		},
	}

	_ = json.NewEncoder(w).Encode(transactions)

}

func createTransactions(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var body, _ = ioutil.ReadAll(r.Body)
	var reqTransactions = Transactions{}

	_ = json.Unmarshal(body, &reqTransactions)

	fmt.Println(reqTransactions)

}
