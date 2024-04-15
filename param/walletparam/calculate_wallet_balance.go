package walletparam

type CalculateWalletBalanceRequest struct {
	User string `json:"user"`
}
type CalculateWalletBalanceResponse struct {
	Balance float64 `json:"balance"`
}
