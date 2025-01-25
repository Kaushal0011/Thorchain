package main

import (
	"fmt"
	"log"
	"project/services"
)

func main() {
	// Prompt the user for a BTC address
	fmt.Print("Enter BTC address: ")
	var btcAddress string
	fmt.Scanln(&btcAddress)

	// Call the service to filter transactions locally
	transactions, err := services.GetCrossChainTransactions(btcAddress, "data/sample.json")
	if err != nil {
		log.Fatalf("Error processing transactions: %v", err)
	}

	// Display the filtered transaction details
	if len(transactions) == 0 {
		fmt.Println("No Ethereum addresses found for the given BTC address.")
	} else {
		fmt.Println("Transaction details:")
		for _, tx := range transactions {
			fmt.Printf("Timestamp: %s, Ethereum Address: %s, Value: %s\n", tx.Timestamp, tx.ETHAddress, tx.Value)
		}
	}
}
