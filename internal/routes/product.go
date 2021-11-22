package routes

import (
"github.com/gofiber/fiber/v2"
"go-fiber-api/internal/handler/http"
"go-fiber-api/pkg/middleware"
)

func ProductV1Route(v1 fiber.Router, hdl *http.Handler) {
	public := v1.Group("/products")
	//private route
	middleware := middleware.NewMiddleware()
	private := middleware.JWTMiddleware(public)
	{
		private.Get("", hdl.GetAllProduct)
		private.Get("/:id", hdl.FindProductById)
		private.Post("", hdl.CreateProduct)
		private.Put("/:id", hdl.UpdateProduct)
		private.Delete("/:id", hdl.DeleteProduct)
	}
}
