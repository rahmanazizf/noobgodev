package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	OrderID      int       `gorm:"primary_key" json:"id"`
	CustomerName string    `json:"customerName"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"-"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime:mili" json:"-"`
	OrderedAt    time.Time `json:"orderedAt"`
	Items        []Item    `gorm:"constraint:OnDelete:CASCADE" json:"items"`
}

type Item struct {
	ItemID      int       `gorm:"primary_key" json:"-"`
	ItemName    string    `json:"name"`
	Description string    `json:"description"`
	Quantity    int       `json:"quantity"`
	OrderID     *uint     `json:"-"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"-"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime:mili" json:"-"`
}

type Connection struct {
	DB *gorm.DB
}
