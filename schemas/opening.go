package schemas

import (
	"gorm.io/gorm"
	"time"
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

type OpeningResponse struct {
	ID        uint           `json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt,omitempty"`
	Role      string         `json:"role"`
	Company   string         `json:"company"`
	Location  string         `json:"location"`
	Remote    bool           `json:"remote"`
	Link      string         `json:"link"`
	Salary    int64          `json:"salary"`
}
