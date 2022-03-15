package main

import (
	"log"
	"net/http"

	"quiz/controllers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/get/wallets", controllers.GetAllWallet).Methods("GET")
	router.HandleFunc("/insert/wallet", controllers.InsertWallet).Methods("POST")
	router.HandleFunc("/delete/wallet/{id_wallet}", controllers.DeleteWallet).Methods("DELETE")
	// router.HandleFunc("/update/user/{user_id}", controllers.UpdateUser).Methods("PUT")

	router.HandleFunc("/get/walletstransactions", controllers.GetAllWalletTransactions).Methods("GET")
	router.HandleFunc("/insert/transaction", controllers.InsertTransaction).Methods("POST")
	// router.HandleFunc("/delete/product/{product_id}", controllers.DeleteProduct).Methods("DELETE")
	// router.HandleFunc("/update/product/{product_id}", controllers.UpdateProduct).Methods("PUT")

	// router.HandleFunc("/get/transactions", controllers.GetAllTransactions).Methods("GET")
	// router.HandleFunc("/insert/transaction", controllers.InsertTransaction).Methods("POST")
	// router.HandleFunc("/delete/transaction/{transaction_id}", controllers.DeleteTransaction).Methods("DELETE")
	// router.HandleFunc("/update/transaction/{transaction_id}", controllers.UpdateTransaction).Methods("PUT")

	// router.HandleFunc("/get/all/{user_id}", controllers.GetAllDetailTransactions).Methods("GET")
	router.HandleFunc("/login/wallet", controllers.Login).Methods("POST")

	http.Handle("/", router)
	log.Println("Connected to port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
