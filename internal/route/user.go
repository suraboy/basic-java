package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/suraboy/go-fiber-api/internal/handler/http"
)

func UserV1Route(v1 fiber.Router, hdl *http.HTTPHandler) {
	r := v1.Group("/v1/users")
	r.Get("", hdl.GetAllUser)
	//r.Get("/:id", http.FindUser)
	//r.Post("", http.CreateUser)
	//r.Put("/:id", http.UpdateUser)
	//r.Delete("/:id", http.DeleteUser)
}
