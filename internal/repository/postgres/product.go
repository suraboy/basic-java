package postgres

import (
	"go-fiber-api/internal/core/domain"
	"gorm.io/gorm/clause"
)

/**
 * Created by boy.sirichai on 22/11/2021 AD
 */
const (
	TableNameProduct = "products"
)

func (r *Postgres) GetProducts(filter domain.Product) ([]domain.Product, error) {
	var model []domain.Product
	result := r.Connection.
		Table(TableNameProduct).
		Order("created_at desc").
		Find(&model, filter)
	return model, result.Error
}


func (r *Postgres) FindProduct(filter domain.Product) (domain.Product, error) {
	var model domain.Product
	result := r.Connection.
		Table(TableNameProduct).
		First(&model, &filter)
	return model, result.Error
}

func (r *Postgres) CreateProduct(filter domain.Product) (domain.Product, error) {
	result := r.Connection.
		Table(TableNameProduct).
		Create(&filter)

	return filter, result.Error
}

func (r *Postgres) UpdateProduct(filter domain.Product, id int) (domain.Product, error) {
	var model domain.Product
	resp := r.Connection.
		Table(TableNameProduct).
		Model(&domain.Product{ID: uint(id)}).
		Clauses(clause.Returning{}).
		Updates(&filter)
	return model, resp.Error
}

func (r *Postgres) DeleteProduct(filter domain.Product) (domain.Product, error) {
	var model domain.Product
	result := r.Connection.
		Table(TableNameProduct).
		Delete(&model, filter)

	return model, result.Error
}