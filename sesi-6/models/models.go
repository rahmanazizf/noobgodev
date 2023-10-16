package models

import "time"

type Products struct {
	ProductID int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Variants struct {
	VariantID   int
	VariantName string
	Qty         int
	ProductID   int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ProductWithVariants struct {
	ProductID   int
	ProductName string
	Variants    []Variants
}
