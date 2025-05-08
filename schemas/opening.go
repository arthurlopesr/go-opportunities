package schemas

import (
	"gorm.io/gorm"
)

type Opening struct {
	gorm.Model
	Role     string `gorm:"type:varchar(100);not null"`
	Company  string `gorm:"type:varchar(100);not null"`
	Location string `gorm:"type:varchar(100)"`
	Remote   bool
	Link     string `gorm:"type:varchar(255)"`
	Salary   int64  `gorm:"not null"`
}
