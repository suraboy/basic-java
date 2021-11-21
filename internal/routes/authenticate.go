package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-api/internal/handler/http"
)

/**
 * Created by boy.sirichai on 21/11/2021 AD
 */


func AuthV1Route(v1 fiber.Router, hdl *http.Handler) {
	r := v1.Group("/oauth")
	r.Post("/login", hdl.Authenticate)
	r.Post("/refresh", hdl.RefreshToken)
	r.Delete("/logout", hdl.Logout)
}