package postgres

import (
	"go-fiber-api/internal/core/domain"
	"gorm.io/gorm/clause"
)

/**
 * Created by boy.sirichai on 8/11/2021 AD
 */
const (
	TableNameUser = "users"
)

func (r *Postgres) GetUsers(filter domain.User) ([]domain.User, error) {
	var model []domain.User
	result := r.Connection.
		Table(TableNameUser).
		Order("created_at desc").
		Find(&model, filter)
	return model, result.Error
}

func (r *Postgres) FindUser(filter domain.User) (domain.User, error) {
	var model domain.User
	result := r.Connection.
		Table(TableNameUser).
		First(&model, &filter)
	return model, result.Error
}

func (r *Postgres) CreateUser(filter domain.User) (domain.User, error) {
	result := r.Connection.
		Table(TableNameUser).
		Create(&filter)

	return filter, result.Error
}

func (r *Postgres) UpdateUser(filter domain.User, id int) (domain.User, error) {
	var model domain.User
	resp := r.Connection.
		Table(TableNameUser).
		Model(&domain.User{ID: uint(id)}).
		Clauses(clause.Returning{}).
		Updates(&filter)
	return model, resp.Error
}

func (r *Postgres) DeleteUser(filter domain.User) (domain.User, error) {
	var model domain.User
	result := r.Connection.
		Table(TableNameUser).
		Delete(&model, filter)

	return model, result.Error
}
