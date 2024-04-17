package discountservice

import (
	"app/adapter/mysql"
	"app/adapter/redis"
	"app/model/discountmodel"
	"app/param/discountparam"
	"app/pkg/errmsg"
	"app/pkg/richerror"
	"errors"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type Service struct {
	db         mysql.Adapter
	redis      redis.Adapter
	walletServ WalletServer
}
type WalletServer interface {
	CheckUsageDiscount(discount string, user string) bool
	CreateTransaction(user string, typeT string, amount float64) error
}

func New(db mysql.Adapter, redis redis.Adapter, walletServ WalletServer) Service {
	return Service{db: db, redis: redis, walletServ: walletServ}
}

func (s Service) ApplyDiscount(ctx context.Context, req discountparam.ApplyDiscountRequest) (discountparam.ApplyDiscountResponse, error) {
	const op = "discountservice.ApplyDiscount"
	var topCode discountmodel.TopUpCode
	res := s.db.Conn().WithContext(ctx).First(&topCode).Where("code = ?", req.DiscountCode)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return discountparam.ApplyDiscountResponse{}, richerror.New(op).WithErr(res.Error).WithKind(richerror.KindInvalid)
	}
	//Start lock
	isUse := s.walletServ.CheckUsageDiscount(req.DiscountCode, req.User)
	if isUse {
		return discountparam.ApplyDiscountResponse{}, richerror.New(op).WithKind(richerror.KindInvalid).WithMessage(errmsg.ErrorMsgAlreadyUsedCode)
	}
	err := s.walletServ.CreateTransaction(req.User, "increase_by_gift", topCode.Amount)
	//End lock
	if err != nil {
		return discountparam.ApplyDiscountResponse{}, richerror.New(op).WithKind(richerror.KindInvalid).WithMessage(errmsg.ErrorMsgAlreadyUsedCode)
	}
	return discountparam.ApplyDiscountResponse{}, nil
}
