package discountvalidator

import (
	"app/model/discountmodel"
	"app/param/discountparam"
	"app/pkg/errmsg"
	"app/pkg/richerror"
	"errors"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gorm.io/gorm"
	"time"
)

func (v Validator) ValidateApplyDiscountRequest(req discountparam.ApplyDiscountRequest) (map[string]string, error) {
	const op = "discountvalidator.ValidateApplyDiscountRequest"
	if err := validation.ValidateStruct(&req,

		validation.Field(&req.User,
			validation.Required,
		),
		validation.Field(&req.DiscountCode,
			validation.Required,
			validation.By(v.checkTopUpCode),
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

func (v Validator) checkTopUpCode(value interface{}) error {
	code := value.(string)
	err := v.db.Conn().Where("code = ?", code).
		Where("? BETWEEN valid_from AND valid_to", time.Now()).
		First(&discountmodel.TopUpCode{}).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("top up code %s not exist", code)
		}
		return fmt.Errorf("top up code %s not exist", code)
	}
	return nil
}
