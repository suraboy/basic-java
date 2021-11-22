package service

import (
	"go-fiber-api/internal/core/domain"
	responseBody "go-fiber-api/pkg/util"
)

/**
 * Created by boy.sirichai on 22/11/2021 AD
 */


func (s Service) GetAllProduct(request domain.Product) ([]domain.Product, error) {
	return s.repository.GetProducts(request)
}


func (s Service) FindProductById(request domain.Product) (domain.Product, error) {
	return s.repository.FindProduct(request)
}

func (s Service) CreateProduct(request domain.Product) (domain.Product, error) {
	res, err := s.repository.CreateProduct(request)

	if err != nil {
		return domain.Product{}, responseBody.SendExceptionError(c, err)
	}
	return res, nil
}

func (s Service) UpdateProductById(request domain.Product, id int) (domain.Product, error) {
	_, err := s.repository.UpdateProduct(request, id)
	if err != nil {
		return domain.Product{}, err
	}
	userInfo, err := s.repository.FindProduct(domain.Product{ID: uint(id)})
	if err != nil {
		return domain.Product{}, err
	}

	return userInfo, nil
}

func (s Service) DestroyProductById(request domain.Product) (domain.Product, error) {
	_, err := s.repository.FindProduct(request)
	if err != nil {
		return domain.Product{}, err
	}

	res, err := s.repository.DeleteProduct(request)
	if err != nil {
		return domain.Product{}, responseBody.SendExceptionError(c, err)
	}
	return res, nil
}