package models

type Products struct {
	ProductID int
	Name      string
	CreatedAt interface{}
	UpdatedAt interface{}
}

type Variants struct {
	VariantID   int
	VariantName string
	Qty         int
	ProductID   int
	CreatedAt   interface{}
	UpdatedAt   interface{}
}

type ProductWithVariants struct {
	ProductID   int
	ProductName string
	Variants    []Variants
}
