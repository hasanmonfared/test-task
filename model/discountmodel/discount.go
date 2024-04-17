package discountmodel

import "time"

type Discount struct {
	ID        uint      `json:"id",gorm:"primary_key"`
	Code      string    `json:"code"`
	Discount  float64   `json:"discount"`
	ValidFrom time.Time `json:"valid_from"`
	ValidTo   time.Time `json:"valid_to"`
}
