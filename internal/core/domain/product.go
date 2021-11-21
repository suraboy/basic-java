package domain

import (
	"time"
)

type Product struct {
	ID   uint `gorm:"primary_key" json:"id"`
	Name []struct {
		En string `json:"en" gorm:"type:varchar(255);"`
		Th string `json:"th" gorm:"type:varchar(255);"`
	}
	Price     float64    `json:"price" validate:"required"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
