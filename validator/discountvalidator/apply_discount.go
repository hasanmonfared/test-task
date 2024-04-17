package discountvalidator

import (
	"app/param/discountparam"
	"app/pkg/errmsg"
	"app/pkg/richerror"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (v Validator) ValidateApplyDiscountRequest(req discountparam.ApplyDiscountRequest) (map[string]string, error) {
	const op = "discountvalidator.ValidateApplyDiscountRequest"
	if err := validation.ValidateStruct(&req,

		validation.Field(&req.User,
			validation.Required,
		),
	); err != nil {
		fieldErrors := make(map[string]string)
		errV, ok := err.(validation.Errors)
		if ok {
			for key, value := range errV {
				if value != nil {
					fieldErrors[key] = value.Error()
				}
			}
		}
		return fieldErrors, richerror.New(op).
			WithMessage(errmsg.ErrorMsgInvalidInput).
			WithKind(richerror.KindInvalid).
			WithMeta(map[string]interface{}{"req": req}).
			WithErr(err)
	}
	return nil, nil
}
