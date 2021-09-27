package models

import (
	"time"
)

type Product struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	Name      Languages  `json:"name,string,omitempty"`
	Price     float64    `json:"price" validate:"required"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type Languages struct {
	En string `json:"en"`
	Th string `json:"th"`
}