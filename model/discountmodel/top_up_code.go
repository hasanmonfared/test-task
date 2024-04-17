package discountmodel

import "time"

type TopUpCode struct {
	ID        uint      `json:"id",gorm:"primary_key"`
	Code      string    `json:"code"`
	Amount    float64   `json:"amount"`
	ValidFrom time.Time `json:"valid_from"`
	ValidTo   time.Time `json:"valid_to"`
}
