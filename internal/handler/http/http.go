package http

import (
	"go-fiber-api/internal/core/port"
	"go-fiber-api/internal/core/service"
	"go-fiber-api/pkg/validators"
)

/**
 * Created by boy.sirichai on 8/11/2021 AD
 */

type HTTPHandler struct {
	svc       port.Service
	validator validators.Validator
}

func NewHTTPHandler(svc *service.Service, validator validators.Validator) *HTTPHandler {
	return &HTTPHandler{
		svc:       svc,
		validator: validator,
	}
}
