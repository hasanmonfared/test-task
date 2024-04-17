package discountmodel

import "time"

type Discount struct {
	ID        uint      `json:"id"`
	Code      string    `json:"code"`
	Discount  float64   `json:"discount"`
	ValidFrom time.Time `json:"valid_from"`
	ValidTo   time.Time `json:"valid_to"`
}
