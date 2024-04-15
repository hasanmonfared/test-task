package walletparam

import "app/model/walletmodel"

type CreateTransactionRequest struct {
	User   string
	Type   walletmodel.Type
	Amount float64
}
type CreateTransactionResponse struct {
}
