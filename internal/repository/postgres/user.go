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

func (r *Postgres) GetUsers(request domain.User) ([]domain.User, error) {
	var model []domain.User
	result := r.Connection.
		Table(TableNameUser).
		Order("created_at desc").
		Find(&model, request)
	return model, result.Error
}

func (r *Postgres) FindUser(request domain.User) (domain.User, error) {
	var model domain.User
	result := r.Connection.
		Table(TableNameUser).
		First(&model, &request)
	return model, result.Error
}

func (r *Postgres) CreateUser(request domain.User) (domain.User, error) {
	result := r.Connection.
		Table(TableNameUser).
		Create(&request)

	return request, result.Error
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

func (r *Postgres) DeleteUser(id int) (domain.User, error) {
	var model domain.User
	result := r.Connection.
		Table(TableNameUser).
		Delete(&model, "id = ?", id)

	return model, result.Error
}
