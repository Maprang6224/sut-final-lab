package entity

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name       string
	Email      string
	CustomerID string
}
