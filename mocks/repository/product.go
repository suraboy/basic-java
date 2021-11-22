package repository

import "go-fiber-api/internal/core/domain"

/**
 * Created by boy.sirichai on 22/11/2021 AD
 */


func (_m *Repository) GetProducts(request domain.Product) ([]domain.Product, error) {
	var _ca []interface{}
	_ca = append(_ca, request)
	ret := _m.Called(_ca...)

	var r0 []domain.Product
	if rf, ok := ret.Get(0).(func(domain.Product) []domain.Product); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Get(0).([]domain.Product)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Product) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindProduct provides a mock function with given fields: model
func (_m *Repository) FindProduct(request domain.Product) (domain.Product, error) {
	var _ca []interface{}
	_ca = append(_ca, request)
	ret := _m.Called(_ca...)

	var r0 domain.Product
	if rf, ok := ret.Get(0).(func(domain.Product) domain.Product); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Get(0).(domain.Product)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Product) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
// CreateProduct provides a mock function with given fields: model
func (_m *Repository) CreateProduct(request domain.Product) (domain.Product, error) {
	ret := _m.Called(request)

	var r0 domain.Product
	if rf, ok := ret.Get(0).(func(domain.Product) domain.Product); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Get(0).(domain.Product)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Product) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Repository) UpdateProduct(request domain.Product, id int) (domain.Product, error) {
	ret := _m.Called(request)

	var r0 domain.Product
	if rf, ok := ret.Get(0).(func(domain.Product) domain.Product); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Get(0).(domain.Product)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Product) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Repository) DeleteProduct(request domain.Product) (domain.Product, error) {
	ret := _m.Called(request)
	var r0 domain.Product
	if rf, ok := ret.Get(0).(func(domain.Product) domain.Product); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Get(0).(domain.Product)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Product) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

