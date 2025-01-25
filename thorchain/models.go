package thorchain

// Coin represents a single coin involved in a transaction
type Coin struct {
	Amount string `json:"amount"`
	Asset  string `json:"asset"`
}

// TransactionInput represents the input details of a transaction
type TransactionInput struct {
	Address string `json:"address"`
	Coins   []Coin `json:"coins"`
	TxID    string `json:"txID"`
}

// TransactionOutput represents the output details of a transaction
type TransactionOutput struct {
	Address   string `json:"address"`
	Affiliate bool   `json:"affiliate,omitempty"`
	Coins     []Coin `json:"coins"`
	Height    string `json:"height"`
	TxID      string `json:"txID"`
}

// Swap represents the swap metadata of a transaction
type Swap struct {
	AffiliateAddress string `json:"affiliateAddress"`
	AffiliateFee     string `json:"affiliateFee"`
	InPriceUSD       string `json:"inPriceUSD"`
	OutPriceUSD      string `json:"outPriceUSD"`
	SwapSlip         string `json:"swapSlip"`
	Memo             string `json:"memo"`
}

// TransactionMetadata represents metadata of a transaction
type TransactionMetadata struct {
	Swap Swap `json:"swap"`
}

// Action represents a single transaction action
type Action struct {
	Date     string              `json:"date"`
	Height   string              `json:"height"`
	In       []TransactionInput  `json:"in"`
	Metadata TransactionMetadata `json:"metadata"`
	Out      []TransactionOutput `json:"out"`
	Pools    []string            `json:"pools"`
	Status   string              `json:"status"`
	Type     string              `json:"type"`
}

// TransactionsResponse represents the structure of the JSON data
type TransactionsResponse struct {
	Actions []Action `json:"actions"`
}
