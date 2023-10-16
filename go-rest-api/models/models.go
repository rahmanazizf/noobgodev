package models

import "time"

type Order struct {
	OrderID      int       `gorm:"primary_key" json:"id"`
	CustomerName string    `json:"customerName"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	OrderedAt    time.Time `json:"orderedAt"`
	Items        []Item    `json:"items"`
}

type Item struct {
	ItemID      int       `gorm:"primary_key" json:"id"`
	ItemName    string    `json:"name"`
	Description string    `json:"description"`
	Quantity    int       `json:"quantity"`
	OrderID     int       `json:"order_id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
