package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-api/internal/handler/http"
	"go-fiber-api/pkg/middleware"
)

func UserV1Route(v1 fiber.Router, hdl *http.HTTPHandler) {
	middleware := middleware.NewMiddleware()
	r := v1.Group("/users")
	r.Get("", middleware.ApiKeyAccess, hdl.GetAllUser)
	r.Get("/:id", middleware.ApiKeyAccess, hdl.FindUserById)
	//r.Post("", hdl.CreateUser)
	//r.Put("/:id", hdl.UpdateUser)
	//r.Delete("/:id", hdl.DeleteUser)
}
