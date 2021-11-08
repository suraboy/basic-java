package postgres

import "github.com/suraboy/go-fiber-api/internal/core/domain/models"

/**
 * Created by boy.sirichai on 8/11/2021 AD
 */


func (r *Postgres) GetAllUser(request *models.User, field ...interface{}) (models.User, error) {
	var model models.User
	resp := r.Connection.
		Table("users").
		Select("id", field...).
		Order("created_at desc").
		First(&model, request)
	return model, resp.Error
}