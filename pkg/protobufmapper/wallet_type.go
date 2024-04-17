package protobufmapper

import (
	"app/contract/goproto/wallet"
	"app/model/walletmodel"
	"app/param/walletparam"
)

func MapProtobufTypeWalletToParam(t wallet.Type) walletmodel.Type {
	switch t {
	case wallet.Type_DEPOSIT:

		return walletmodel.Deposit
	case wallet.Type_IncreaseByGift:
		return walletmodel.IncreaseByGift
	case wallet.Type_WITHDRAWAL:

		return walletmodel.Withdrawal
	}
	return walletmodel.Unknown
}

func MapTypeWalletParamToProtobuf(t walletmodel.Type) wallet.Type {
	switch t {
	case walletmodel.Deposit:
		return wallet.Type_DEPOSIT
	case walletmodel.IncreaseByGift:
		return wallet.Type_IncreaseByGift
	case walletmodel.Withdrawal:
		return wallet.Type_WITHDRAWAL
	}
	return wallet.Type_Unknown
}
func MapResponseParamToProtobuf(res walletparam.CreateTransactionResponse) *wallet.CreateTransactionResponse {
	return &wallet.CreateTransactionResponse{}
}
func MapProtobufToResponseParam(res *wallet.CreateTransactionResponse) walletparam.CreateTransactionResponse {
	return walletparam.CreateTransactionResponse{}
}
func MapTypeStringToProtobuf(s string) wallet.Type {
	switch s {
	case "deposit":
		return wallet.Type_DEPOSIT
	case "increase_by_gift":
		return wallet.Type_IncreaseByGift
	case "withdrawal":
		return wallet.Type_WITHDRAWAL
	}
	return wallet.Type_Unknown
}
