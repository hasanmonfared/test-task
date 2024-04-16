package walletservice

import (
	"app/adapter/mysql"
	"app/model/walletmodel"
	"app/param/walletparam"
	"app/pkg/richerror"
	"golang.org/x/net/context"
)

type Service struct {
	db mysql.Adapter
}

func New(db mysql.Adapter) Service {
	return Service{db: db}
}
func (s Service) CreateTransaction(context context.Context, req walletparam.CreateTransactionRequest) (walletparam.CreateTransactionResponse, error) {
	const op = "walletservice.createTransaction"
	transaction := walletmodel.Transaction{
		User:   req.User,
		Type:   req.Type,
		Amount: req.Amount,
	}
	err := s.db.Conn().WithContext(context).Create(&transaction).Error
	if err != nil {
		richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}
	return walletparam.CreateTransactionResponse{}, nil
}
func (s Service) CalculateWalletBalance(ctx context.Context, req walletparam.CalculateWalletBalanceRequest) (walletparam.CalculateWalletBalanceResponse, error) {
	const op = "walletservice.calculateWalletBalance"
	var balance float64
	var transactions []walletmodel.Transaction
	s.db.Conn().Where("user = ?", req.User).Find(&transactions)

	for _, transaction := range transactions {
		if transaction.Type == walletmodel.Deposit {
			balance += transaction.Amount
		} else if transaction.Type == walletmodel.Withdrawal {
			balance -= transaction.Amount
		}
	}
	return walletparam.CalculateWalletBalanceResponse{Balance: balance}, nil
}
