package walletmodel

type Type string

const (
	Deposit        Type = "deposit"
	Withdrawal          = "withdrawal"
	IncreaseByGift      = "increase_by_gift"
	Unknown             = "unknown"
)
