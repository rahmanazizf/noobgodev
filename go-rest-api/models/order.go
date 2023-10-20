package models

import "time"

type Order struct {
	OrderID      int       `gorm:"primary_key" json:"id"`
	CustomerName string    `json:"customerName"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"-"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime:mili" json:"-"`
	OrderedAt    time.Time `json:"orderedAt"`
	Items        []Item    `gorm:"constraint:OnDelete:CASCADE" json:"items"`
}
