package models

import (
	"gorm.io/gorm"
)

type Sales struct {
	gorm.Model
	Amount		float64 `json:"Amount"`
}

type Payments struct {
	gorm.Model
	Amount		float64	 `json:"amount"`
	SupplierID	uint	 `json:"supplierID"`
	Supplier	Supplier `json:"Supplier" gorm:"foreignKey:SupplierID"`
}

type Debt struct {
	gorm.Model
	Amount		float64
	Description	string 	 `json:"Description"`
	SupplierID	uint	 `json:"SupplierID"`
	Supplier	Supplier `json:"supplier" gorm:"foreignKey:SupplierID"`
}

type Supplier struct {
	gorm.Model
	Name 		string		`json:"Name" gorm:"unique"`
}