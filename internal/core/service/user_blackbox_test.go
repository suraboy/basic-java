package service_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"

	"go-fiber-api/internal/core/domain"
	"go-fiber-api/internal/core/service"
)

/**
 * Created by boy.sirichai on 22/11/2021 AD
 */

func TestGetUser(t *testing.T) {
	cfg := service.Config{
		Repository: mocker.Repository,
	}
	svc := service.New(cfg)
	defaultRequest := domain.User{}
	t.Run("GetUser() doesn't got any error should get succeed", func(t *testing.T) {
		mocker.Repository.
			On(
				method.Repository.GetUsers,
				mock.AnythingOfType("domain.User")).
			Return([]domain.User{}, nil).
			Once()
		_, err := svc.GetAllUser(defaultRequest)
		assert.NoError(t, err)
		mocker.Repository.AssertExpectations(t)
	})
}

func TestGetUserById(t *testing.T) {
	cfg := service.Config{
		Repository: mocker.Repository,
	}
	svc := service.New(cfg)
	defaultRequest := domain.User{}
	t.Run("FindUserById() doesn't got any error should get succeed", func(t *testing.T) {
		mocker.Repository.
			On(
				method.Repository.FindUser,
				mock.AnythingOfType("domain.User")).
			Return(domain.User{}, nil).
			Once()
		_, err := svc.FindUserById(defaultRequest)
		assert.NoError(t, err)
		mocker.Repository.AssertExpectations(t)
	})
}

func TestCreateUser(t *testing.T) {
	cfg := service.Config{
		Repository: mocker.Repository,
	}
	svc := service.New(cfg)
	defaultRequest := domain.User{}
	t.Run("CreateUser() doesn't got any error should get succeed", func(t *testing.T) {
		mocker.Repository.
			On(
				method.Repository.CreateUser,
				mock.AnythingOfType("domain.User"),
			).
			Return(domain.User{}, nil).
			Once()
		_, err := svc.CreateUser(defaultRequest)
		assert.NoError(t, err)
		mocker.Repository.AssertExpectations(t)
	})
}


func TestUpdateUser(t *testing.T) {
	cfg := service.Config{
		Repository: mocker.Repository,
	}
	svc := service.New(cfg)
	defaultRequest := domain.User{}
	t.Run("UpdateUser() doesn't got any error should get succeed", func(t *testing.T) {
		mocker.Repository.
			On(
				method.Repository.UpdateUser,
				mock.AnythingOfType("domain.User"),
			).
			Return(domain.User{}, nil).
			Once()
		mocker.Repository.
			On(
				method.Repository.FindUser,
				mock.AnythingOfType("domain.User"),
			).
			Return(domain.User{}, nil).
			Once()
		_, err := svc.UpdateUserById(defaultRequest,1)
		assert.NoError(t, err)
		mocker.Repository.AssertExpectations(t)
	})
}


func TestDeleteUser(t *testing.T) {
	cfg := service.Config{
		Repository: mocker.Repository,
	}
	svc := service.New(cfg)
	defaultRequest := domain.User{}
	t.Run("DeleteUser() doesn't got any error should get succeed", func(t *testing.T) {
		mocker.Repository.
			On(
				method.Repository.FindUser,
				mock.AnythingOfType("domain.User"),
			).
			Return(domain.User{}, nil).
			Once()
		mocker.Repository.
			On(
				method.Repository.DeleteUser,
				mock.AnythingOfType("domain.User"),
			).
			Return(domain.User{}, nil).
			Once()
		_, err := svc.DestroyUserById(defaultRequest)
		assert.NoError(t, err)
		mocker.Repository.AssertExpectations(t)
	})
}
