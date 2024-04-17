package discountvalidator

import "app/adapter/mysql"

type Validator struct {
	db mysql.Adapter
}

func New(db mysql.Adapter) Validator {
	return Validator{
		db: db,
	}
}
