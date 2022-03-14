package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	// "log"
	// "net/http"
	// "strconv"
	"quiz/model"
	//"github.com/gorilla/mux"
)

func GetAllWalletTransactions(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	query := "SELECT * FROM table_wallet JOIN table_transaction ON table_wallet.id_wallet = table_transaction.idWallet"

	rows, err := db.Query(query)

	if err != nil {
		var response model.ErrorResponse
		response.Status = "500"
		response.Message = "The server encountered an unexpected condition that prevented it from fulfilling the request."
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	var transaction model.Transaction
	var wallet model.Wallet
	var transactions []model.Transaction
	var wallets []model.Wallet

	for rows.Next() {
		if err := rows.Scan(&wallet.Id_Wallet, &wallet.Username, &wallet.Password, &wallet.Currency, &wallet.DisableUser, &transaction.IdTransaction, &transaction.IdWallet, &transaction.Datetime, &transaction.Amount, &transaction.Description); err != nil {
			var response model.ErrorResponse
			response.Status = "500"
			response.Message = "The server encountered an unexpected condition that prevented it from fulfilling the request."
			log.Fatal(err.Error())
		} else {
			transactions = append(transactions, transaction)
			wallets = append(wallets, wallet)
		}
	}

	var response model.Response

	response.Status = "200"
	response.Message = "Indicates that the request has succeeded."
	response.DataTransaction = transactions
	response.DataWallet = wallets

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
