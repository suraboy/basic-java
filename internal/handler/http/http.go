package http

import (
	"go-fiber-api/internal/core/port"
	"go-fiber-api/pkg/validators"
)

/**
 * Created by boy.sirichai on 8/11/2021 AD
 */

// Handler ...
type Handler struct {
	svc       port.Service
	validator validators.Validator
}

// NewHandler ...
// application HTTP handler
func NewHandler(svc port.Service, validator validators.Validator) *Handler {
	return &Handler{
		svc:       svc,
		validator: validator,
	}
}
