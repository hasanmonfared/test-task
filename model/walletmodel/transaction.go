package walletmodel

import "time"

type Transaction struct {
	ID        uint      `json:"id",gorm:"primary_key"`
	User      string    `json:"user_id"`
	Type      Type      `json:"type",gorm:"type:enum('deposit', 'withdrawal','unknown')"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}
