package validators_test

import (
	"go-fiber-api/pkg/validators"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testSubject validators.Validator
)

func TestMain(m *testing.M) {
	testSubject = validators.NewValidator()
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestValidateStruct(t *testing.T) {
	var err error

	t.Run("ValidateStruct(), Input a struct followed by a valid rule of value, Should get succeed", func(t *testing.T) {
		mockInput := struct {
			// validate is that message uppercase
			Message string `validate:"uppercase"`
		}{
			Message: "HELLO",
		}
		err = testSubject.ValidateStruct(mockInput)
		assert.NoError(t, err)
	})
	t.Run("ValidateStruct(), Input a struct followed by a invalid rule of value, Should get an error", func(t *testing.T) {
		mockInput := struct {
			// validate is that message uppercase
			Message string `validate:"uppercase"`
		}{
			Message: "hello",
		}
		err = testSubject.ValidateStruct(mockInput)
		assert.Error(t, err)
	})
}
