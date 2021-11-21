package postgres

import (
	"go-fiber-api/internal/core/domain"
)

/**
 * Created by boy.sirichai on 8/11/2021 AD
 */

func (r *Postgres) GetUsers(request domain.User) ([]domain.User, error) {
	var model []domain.User
	result := r.Connection.
		Table("users").
		Order("created_at desc").
		Find(&model, request)
	return model, result.Error
}

func (r *Postgres) FindUser(request domain.User) (domain.User, error) {
	var model domain.User
	result := r.Connection.
		Table("users").
		First(&model, &request)
	return model, result.Error
}

func (r *Postgres) CreateUser(request domain.User) (domain.User, error) {
	result := r.Connection.
		Table("users").
		Create(&request)

	return request, result.Error
}

func (r *Postgres) UpdateUser(request domain.User) (domain.User, error) {
	var model domain.User
	result := r.Connection.
		Table("users").
		Where(&model).
		Updates(request)

	return model, result.Error
}

func (r *Postgres) DeleteUser(request domain.User, id string) (domain.User, error) {
	var model domain.User
	result := r.Connection.
		Table("users").
		Delete(&request, id)

	return model, result.Error
}
