package postgres

import (
	"go-fiber-api/internal/core/domain/models"
)

/**
 * Created by boy.sirichai on 8/11/2021 AD
 */

func (r *Postgres) GetUsers(request *models.User) ([]models.User, error) {
	var model []models.User
	response := r.Connection.
		Table("users").
		Order("created_at desc").
		Find(&model, request)
	return model, response.Error
}

func (r *Postgres) FindUser(request *models.User) (models.User, error) {
	var model models.User
	response := r.Connection.
		Table("users").
		First(&model, &request)
	return model, response.Error
}
