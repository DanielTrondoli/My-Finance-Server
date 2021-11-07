package transaction

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/DanielTrondoli/My-Finance-Server/model/transaction"
	"github.com/DanielTrondoli/My-Finance-Server/util"
)

// GetTransaction blablabal
func GetTransaction(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-type", "application/json")

	dateRecived := util.StringToTime("1992-11-07T16:52:00")

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

// CreateTransactions blablabal
func CreateTransactions(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var body, _ = ioutil.ReadAll(r.Body)
	var reqTransactions = transaction.Transactions{}

	_ = json.Unmarshal(body, &reqTransactions)

	fmt.Println(reqTransactions)

}
