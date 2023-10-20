package models

import "time"

type Item struct {
	ItemID      int       `gorm:"primary_key" json:"-"`
	ItemName    string    `json:"name"`
	Description string    `json:"description"`
	Quantity    int       `json:"quantity"`
	OrderID     *uint     `json:"-"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"-"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime:mili" json:"-"`
}
