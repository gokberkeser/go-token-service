package main

import (
	"example.com/go-token-service/src"
	"net/http"
)

func main() {
	print("Starting server on port 8080")
	http.HandleFunc("/transfer", src.TransferToken)
	http.HandleFunc("/balance", src.BalanceHandler)

	http.ListenAndServe(":8080", nil)
}
