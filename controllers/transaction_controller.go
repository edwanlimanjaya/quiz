package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"quiz/model"
	"strconv"
)

func GetAllWalletTransactions(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	query := "SELECT * FROM table_wallet JOIN table_transaction ON table_wallet.id_wallet = table_transaction.idWallet"

	rows, err := db.Query(query)

	if err != nil {
		log.Fatal(err.Error())
	}

	var transaction model.Transaction
	var wallet model.Wallet
	var transactions []model.Transaction
	var wallets []model.Wallet

	for rows.Next() {
		if err := rows.Scan(&wallet.Id_Wallet, &wallet.Username, &wallet.Password, &wallet.Currency, &wallet.DisableUser, &transaction.IdTransaction, &transaction.IdWallet, &transaction.Datetime, &transaction.Amount, &transaction.Description); err != nil {
			var response model.Response
			response.Status = "500"
			response.Message = "The server encountered an unexpected condition that prevented it from fulfilling the request."
			log.Fatal(err.Error())
		} else {
			transactions = append(transactions, transaction)
			wallets = append(wallets, wallet)
		}
	}

	var response model.ResponseAll

	response.Status = "200"
	response.Message = "Indicates that the request has succeeded."
	response.DataTransaction = transactions
	response.DataWallet = wallets

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func InsertTransaction(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	err := r.ParseForm()

	if err != nil {
		return
	}

	idWallet, _ := strconv.Atoi(r.Form.Get("idWallet"))
	amount, _ := strconv.Atoi(r.Form.Get("amount"))
	description := r.Form.Get("description")

	query := "Insert into table_transaction (idWallet, amount, description) values (?,?,?)"

	_, errQuery := db.Exec(query, idWallet, amount, description)

	if errQuery != nil {
		log.Fatal(err.Error())
	}

	var response model.Response
	response.Status = "200"
	response.Message = "Indicates that the request has succeeded."

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
