package walletparam

import "app/model/walletmodel"

type CreateTransactionRequest struct {
	User   string           `json:"user"`
	Type   walletmodel.Type `json:"type"`
	Amount float64          `json:"amount"`
}
type CreateTransactionResponse struct {
}
