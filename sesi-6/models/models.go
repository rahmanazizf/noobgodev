package models

type Products struct {
	Name string
}

type Variants struct {
	VariantName string
	Qty         int
	ProductID   int
}
