package main

import (
	"app/adapter/walletadapter"
	"app/model/walletmodel"
	"app/param/walletparam"
	"context"
	"log"
)

func main() {
	client := walletadapter.New("localhost:9898")
	_, err := client.CreateTransaction(context.Background(), walletparam.CreateTransactionRequest{
		User:   "09375",
		Type:   walletmodel.Deposit,
		Amount: 100,
	})
	if err != nil {
		log.Fatalf("CreateTransaction error: %v", err)
	}

}
