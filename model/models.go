package model

type Wallet struct {
	Id_Wallet   int    `json : id_wallet`
	Currency    string `json : currency`
	Username    string `json : username`
	Password    string `json : password`
	DisableUser bool   `json : disableUser`
}

type Transaction struct {
	IdTransaction int    `json : idTransaction`
	IdWallet      int    `json : idWallet`
	Datetime      string `json : datetime`
	Amount        int    `json : amount`
	Description   string `json : description`
}

type WalletResponse struct {
	Status  string `json : status`
	Message string `json : message`
	Data    Wallet `json : data`
}

type WalletsResponse struct {
	Status  string   `json : status`
	Message string   `json : message`
	Data    []Wallet `json : data`
}

type TransactionResponse struct {
	Status  string      `json : status`
	Message string      `json : message`
	Data    Transaction `json : data`
}

type TransactionsResponse struct {
	Status  string        `json : status`
	Message string        `json : message`
	Data    []Transaction `json : data`
}

type ResponseAll struct {
	Status          string        `json : status`
	Message         string        `json : message`
	DataTransaction []Transaction `json : data transaction`
	DataWallet      []Wallet      `json : data wallet`
}

type Response struct {
	Status  string `json : status`
	Message string `json : message`
}
