package walletservice

import (
	"app/model/walletmodel"
	"app/param/walletparam"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func New(db *gorm.DB) Service {
	return Service{db: db}
}
func (s Service) CreateTransaction(context context.Context, req walletparam.CreateTransactionRequest) (walletparam.CreateTransactionResponse, error) {
	transaction := walletmodel.Transaction{
		User:   req.User,
		Type:   req.Type,
		Amount: req.Amount,
	}
	s.db.Create(&transaction)
	return walletparam.CreateTransactionResponse{}, nil
}
