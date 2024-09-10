package routes

import (
	"biodata/handler"
	"biodata/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

func RouteHandler(r *fiber.App) {

r.Get("/app/v1/users",  handler.BiodataHandler)
r.Post("/app/v1/create/user",middleware.AuthLimiter, handler.BiodataHandlerCreate)
r.Get("/app/v1/user/:id", middleware.AuthMiddleware, handler.UserHandlerGetById)
r.Put("/app/v1/user/:id", handler.UserHandlerUpdate)
r.Put("/app/v1/user/:id/update-email", handler.UserHandlerUpdateEmail)
r.Delete("/app/v1/user/:id", handler.UserHandlerDelete)
r.Post("/app/v1/login",middleware.AuthLimiter, handler.LoginHandler)

r.All("/app/v1/service/*", middleware.AuthMiddleware, func(c *fiber.Ctx) error  {
	return proxy.Do(c, "http://localhost:7001" + c.OriginalURL())
})
}