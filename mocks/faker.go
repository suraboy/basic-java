package mocks

import (
	"errors"
	"go-fiber-api/mocks/repository"
)

var (
	// ErrFoo ... is fake error, it can use globally in the application
	ErrFoo = errors.New("error faker")

	// Mocker ... is an exporter to export mock interfaces data structure
	Mocker = &mockFactory{
		Repository: new(repository.Repository),
	}

	// Method ... is an exporter to export methods to use when calling a mock interface
	Method = struct {
		Repository struct {
			GetUsers,
			FindUser,
			CreateUser,
			UpdateUser,
			DeleteUser,
			GetProducts,
			FindProduct,
			CreateProduct,
			UpdateProduct,
			DeleteProduct string
		}
	}{
		Repository: struct {
			GetUsers,
			FindUser,
			CreateUser,
			UpdateUser,
			DeleteUser,
			GetProducts,
			FindProduct,
			CreateProduct,
			UpdateProduct,
			DeleteProduct string
		}{
			"GetUsers",
			"FindUser",
			"CreateUser",
			"UpdateUser",
			"DeleteUser",
			"GetProducts",
			"FindProduct",
			"CreateProduct",
			"UpdateProduct",
			"DeleteProduct",
		},
	}
)

// MockFactory is a data structure that holds all of the mock interfaces
type mockFactory struct {
	Repository *repository.Repository
}
