package discountservice

import (
	"app/adapter/mysql"
	"app/adapter/redis"
	"app/param/discountparam"
	"golang.org/x/net/context"
)

type Service struct {
	db    mysql.Adapter
	redis redis.Adapter
}

func New(db mysql.Adapter, redis redis.Adapter) Service {
	return Service{db: db, redis: redis}
}

func (s Service) ApplyDiscount(ctx context.Context, req discountparam.ApplyDiscountRequest) (discountparam.ApplyDiscountResponse, error) {

	panic("implement me")
	//Check exists code(validation)

	// lock request
	// create queue with redis

}
