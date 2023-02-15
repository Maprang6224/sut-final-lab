package entity

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name       string `valid:"required~กรุณากรอกชื่อ"`
	Email      string
	CustomerID string `valid:"required~CustermerID not match, matches((L|M|H)(1|2|3|4|5|6|7)(1|2|3|4|5|6|7)(1|2|3|4|5|6|7)(1|2|3|4|5|6|7)(1|2|3|4|5|6|7)(1|2|3|4|5|6|7)(1|2|3|4|5|6|7)$)~CustermerID not match"`
}
