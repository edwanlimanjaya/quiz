package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	// "log"
	// "net/http"
	// "strconv"
	"quiz/model"

	"github.com/gorilla/mux"
	//"github.com/gorilla/mux"
)

func GetAllWallet(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	query := "Select * from table_wallet"

	rows, err := db.Query(query)

	fmt.Println(err)

	if err != nil {
		var response model.Response
		response.Status = "500"
		response.Message = "The server encountered an unexpected condition that prevented it from fulfilling the request."
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	var wallet model.Wallet
	var wallets []model.Wallet

	for rows.Next() {
		if err != rows.Scan(&wallet.Id_Wallet, &wallet.Currency, &wallet.Username, &wallet.Password, &wallet.DisableUser) {
			var response model.Response
			response.Status = "404"
			response.Message = "The server can not find the requested resource."
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
			return
		} else {
			wallets = append(wallets, wallet)
		}
	}

	var response model.WalletsResponse

	response.Status = "200"
	response.Message = "Indicates that the request has succeeded."
	response.Data = wallets

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func InsertWallet(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		log.Fatal(err.Error())
	}

	currency := r.Form.Get("currency")
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	disableUser, _ := strconv.ParseBool(r.Form.Get("disableUser"))

	query := "Insert into table_wallet (currency, username, password, disableUser) Values(?,?,?,?)"

	result, errQuery := db.Exec(query, currency, username, password, disableUser)
	fmt.Println(errQuery)
	var wallet model.Wallet
	if errQuery != nil {
		var response model.Response
		response.Status = "500"
		response.Message = "The server encountered an unexpected condition that prevented it from fulfilling the request."
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	} else {
		temp, _ := result.LastInsertId()
		wallet.Id_Wallet = int(temp)
		wallet.Currency = currency
		wallet.Username = username
		wallet.Password = password
	}

	var response model.WalletResponse
	response.Status = "200"
	response.Message = "Indicates that the request has succeeded."
	response.Data = wallet

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func DeleteWallet(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		log.Fatal(err.Error())
	}

	vars := mux.Vars(r)
	Id_Wallet := vars["id_wallet"]

	query := "UPDATE table_wallet SET disableUser = 1 WHERE id_wallet = ? and disableUser = false"

	result, errQuery := db.Exec(query, Id_Wallet)

	temp, _ := result.RowsAffected()

	if int(temp) <= 0 {
		if errQuery != nil {
			log.Fatal(err.Error())
		}
		var response model.Response
		response.Status = "404"
		response.Message = "The server can not find the requested resource."
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	var response model.WalletResponse
	response.Status = "204"
	response.Message = "The server has fulfilled the request but does not need to return a response body. The server may return the updated meta information."

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func UpdateWallet(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	err := r.ParseForm()

	if err != nil {
		return
	}

	vars := mux.Vars(r)
	id_wallet := vars["id_wallet"]

	username := r.Form.Get("username")
	password := r.Form.Get("password")
	currency := r.Form.Get("currency")

	query := "Update table_wallet set username = ?, password = ?, currency = ? where id_wallet = ?"

	result, errQuery := db.Exec(query, username, password, currency, id_wallet)

	temp, _ := result.RowsAffected()

	if int(temp) <= 0 {
		if errQuery != nil {
			log.Fatal(err.Error())
		}
		var response model.Response
		response.Status = "404"
		response.Message = "The server can not find the requested resource."
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	var response model.Response
	response.Status = "204"
	response.Message = "The server has fulfilled the request but does not need to return a response body. The server may return the updated meta information."

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func Login(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	err := r.ParseForm()

	if err != nil {
		log.Fatal(err.Error())
	}

	username := r.Form.Get("username")
	password := r.Form.Get("password")

	query := "Select username, password from table_wallet where username = ? and password = ?"

	var wallet model.Wallet

	errQuery := db.QueryRow(query, username, password).Scan(&wallet.Username, &wallet.Password)

	if errQuery != nil {
		var response model.Response
		response.Status = "404"
		response.Message = "The server can not find the requested resource."
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	var response model.Response
	response.Status = "204"
	response.Message = "The server has fulfilled the request but does not need to return a response body. The server may return the updated meta information."

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
