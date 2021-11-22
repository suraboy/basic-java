package service_test

import (
	"go-fiber-api/mocks"
	"os"
	"testing"
)

var (
	mocker = mocks.Mocker
	method = mocks.Method
)

func TestMain(m *testing.M) {
	exitCode := m.Run()
	os.Exit(exitCode)
}
