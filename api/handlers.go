package api

import (
	"encoding/json"
	"net/http"
	"project/services"
)

// Response defines the structure of the API response
type Response struct {
	EthereumAddresses []string `json:"ethereum_addresses"`
	Error             string   `json:"error,omitempty"`
}

// CrossChainHandler processes requests to fetch ETH addresses for a BTC address
func CrossChainHandler(w http.ResponseWriter, r *http.Request) {
	// Get the BTC address from query parameters
	btcAddress := r.URL.Query().Get("btc_address")
	if btcAddress == "" {
		http.Error(w, "Missing btc_address query parameter", http.StatusBadRequest)
		return
	}

	// Fetch the ETH addresses
	ethAddresses, err := services.GetCrossChainTransactions(btcAddress, "http://localhost:8081/sample.json")
	if err != nil {
		response := Response{Error: err.Error()}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Respond with the Ethereum addresses
	response := Response{EthereumAddresses: ethAddresses}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
