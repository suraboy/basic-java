package repository

import "go-fiber-api/internal/core/domain"

/**
 * Created by boy.sirichai on 22/11/2021 AD
 */


func (_m *Repository) GetUsers(request domain.User) ([]domain.User, error) {
	var _ca []interface{}
	_ca = append(_ca, request)
	ret := _m.Called(_ca...)

	var r0 []domain.User
	if rf, ok := ret.Get(0).(func(domain.User) []domain.User); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Get(0).([]domain.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.User) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindUser provides a mock function with given fields: model
func (_m *Repository) FindUser(request domain.User) (domain.User, error) {
	var _ca []interface{}
	_ca = append(_ca, request)
	ret := _m.Called(_ca...)

	var r0 domain.User
	if rf, ok := ret.Get(0).(func(domain.User) domain.User); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.User) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
// CreateUser provides a mock function with given fields: model
func (_m *Repository) CreateUser(request domain.User) (domain.User, error) {
	ret := _m.Called(request)

	var r0 domain.User
	if rf, ok := ret.Get(0).(func(domain.User) domain.User); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.User) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Repository) UpdateUser(request domain.User, id int) (domain.User, error) {
	ret := _m.Called(request)

	var r0 domain.User
	if rf, ok := ret.Get(0).(func(domain.User) domain.User); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.User) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Repository) DeleteUser(request domain.User) (domain.User, error) {
	ret := _m.Called(request)
	var r0 domain.User
	if rf, ok := ret.Get(0).(func(domain.User) domain.User); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.User) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
