package models

import "time"

type Order struct {
	OrderID      int       `gorm:"primary_key" json:"id"`
	CustomerName string    `json:"customerName"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
	OrderedAt    time.Time `json:"orderedAt"`
	Items        []Item    `json:"items"`
}

type Item struct {
	ItemID      int       `gorm:"primary_key" json:"-"`
	ItemName    string    `json:"name"`
	Description string    `json:"description"`
	Quantity    int       `json:"quantity"`
	OrderID     uint      `json:"-"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}
