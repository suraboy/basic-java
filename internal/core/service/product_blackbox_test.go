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

func TestGetProduct(t *testing.T) {
	cfg := service.Config{
		Repository: mocker.Repository,
	}
	svc := service.New(cfg)
	defaultRequest := domain.Product{}
	t.Run("GetProduct() doesn't got any error should get succeed", func(t *testing.T) {
		mocker.Repository.
			On(
				method.Repository.GetProducts,
				mock.AnythingOfType("domain.Product")).
			Return([]domain.Product{}, nil).
			Once()
		_, err := svc.GetAllProduct(defaultRequest)
		assert.NoError(t, err)
		mocker.Repository.AssertExpectations(t)
	})
}

func TestGetProductById(t *testing.T) {
	cfg := service.Config{
		Repository: mocker.Repository,
	}
	svc := service.New(cfg)
	defaultRequest := domain.Product{}
	t.Run("FindProductById() doesn't got any error should get succeed", func(t *testing.T) {
		mocker.Repository.
			On(
				method.Repository.FindProduct,
				mock.AnythingOfType("domain.Product")).
			Return(domain.Product{}, nil).
			Once()
		_, err := svc.FindProductById(defaultRequest)
		assert.NoError(t, err)
		mocker.Repository.AssertExpectations(t)
	})
}

func TestCreateProduct(t *testing.T) {
	cfg := service.Config{
		Repository: mocker.Repository,
	}
	svc := service.New(cfg)
	defaultRequest := domain.Product{}
	t.Run("CreateProduct() doesn't got any error should get succeed", func(t *testing.T) {
		mocker.Repository.
			On(
				method.Repository.CreateProduct,
				mock.AnythingOfType("domain.Product"),
			).
			Return(domain.Product{}, nil).
			Once()
		_, err := svc.CreateProduct(defaultRequest)
		assert.NoError(t, err)
		mocker.Repository.AssertExpectations(t)
	})
}


func TestUpdateProduct(t *testing.T) {
	cfg := service.Config{
		Repository: mocker.Repository,
	}
	svc := service.New(cfg)
	defaultRequest := domain.Product{}
	t.Run("UpdateProduct() doesn't got any error should get succeed", func(t *testing.T) {
		mocker.Repository.
			On(
				method.Repository.UpdateProduct,
				mock.AnythingOfType("domain.Product"),
			).
			Return(domain.Product{}, nil).
			Once()
		mocker.Repository.
			On(
				method.Repository.FindProduct,
				mock.AnythingOfType("domain.Product"),
			).
			Return(domain.Product{}, nil).
			Once()
		_, err := svc.UpdateProductById(defaultRequest,1)
		assert.NoError(t, err)
		mocker.Repository.AssertExpectations(t)
	})
}


func TestDeleteProduct(t *testing.T) {
	cfg := service.Config{
		Repository: mocker.Repository,
	}
	svc := service.New(cfg)
	defaultRequest := domain.Product{}
	t.Run("DeleteProduct() doesn't got any error should get succeed", func(t *testing.T) {
		mocker.Repository.
			On(
				method.Repository.FindProduct,
				mock.AnythingOfType("domain.Product"),
			).
			Return(domain.Product{}, nil).
			Once()
		mocker.Repository.
			On(
				method.Repository.DeleteProduct,
				mock.AnythingOfType("domain.Product"),
			).
			Return(domain.Product{}, nil).
			Once()
		_, err := svc.DestroyProductById(defaultRequest)
		assert.NoError(t, err)
		mocker.Repository.AssertExpectations(t)
	})
}
