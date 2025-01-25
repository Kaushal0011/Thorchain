package services

import (
	"encoding/json"
	"fmt"
	"project/thorchain"
	"project/utils"
	"strings"
)

// TransactionDetails holds information about a cross-chain transaction
type TransactionDetails struct {
	Timestamp  string
	ETHAddress string
	Value      string
}

// GetCrossChainTransactions filters the JSON data for a given BTC address
func GetCrossChainTransactions(btcAddress, filePath string) ([]TransactionDetails, error) {
	// Read data from the JSON file
	data, err := utils.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	// Parse the JSON into the TransactionsResponse model
	var transactions thorchain.TransactionsResponse
	if err := json.Unmarshal(data, &transactions); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	// Filter transactions by the BTC address
	var details []TransactionDetails
	for _, action := range transactions.Actions {
		for _, input := range action.In {
			// Check if the BTC address matches
			if input.Address == btcAddress {
				// Extract the Ethereum address from metadata.swap.memo
				memo := action.Metadata.Swap.Memo
				if strings.Contains(memo, "=:ETH.ETH:") {
					parts := strings.Split(memo, "=:ETH.ETH:")
					if len(parts) > 1 {
						ethAddress := strings.Split(parts[1], ":")[0] // Extract ETH address

						// Collect transaction details
						details = append(details, TransactionDetails{
							Timestamp:  action.Date,
							ETHAddress: ethAddress,
							Value:      input.Coins[0].Amount,
						})
					}
				}
			}
		}
	}

	return details, nil
}
