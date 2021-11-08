package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/suraboy/go-fiber-api/internal/handler/http"
)

func UserV1Route(v1 fiber.Router, hdl *http.HTTPHandler) {
	r := v1.Group("/users")
	r.Get("", hdl.GetAllUser)
	r.Get("/:id", hdl.FindUserById)
	//r.Post("", hdl.CreateUser)
	//r.Put("/:id", hdl.UpdateUser)
	//r.Delete("/:id", hdl.DeleteUser)
}
