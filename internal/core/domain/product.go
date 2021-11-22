package domain

import (
	"time"
)

type Product struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	Name      string   `json:"name" validate:"required"`
	Price     float64    `json:"price" validate:"required"`
	CreatedBy string     `json:"created_by"`
	UpdatedBy string     `json:"updated_by"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
