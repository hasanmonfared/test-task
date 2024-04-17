package discounthandler

import (
	"app/service/discountservice"
	"app/validator/discountvalidator"
)

type Handler struct {
	svc        discountservice.Service
	validation discountvalidator.Validator
}

func New(svc discountservice.Service, validation discountvalidator.Validator) Handler {
	return Handler{
		svc:        svc,
		validation: validation,
	}
}
