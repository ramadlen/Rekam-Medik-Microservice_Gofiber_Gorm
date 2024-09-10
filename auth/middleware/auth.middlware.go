package middleware

import (
	"biodata/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)
func AuthMiddleware(ctx *fiber.Ctx) error {
token := ctx.Get("x-token")
if token == "" {
	return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"messsage": "unathenthicated",
	})
}
_, err  := utils.VerifyToken(token)
if err != nil{
	return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"message": "unauthenticated",
		})
	}
	return ctx.Next()
}


func PermissionCreate(ctx *fiber.Ctx) error {
	return ctx.Next()
}


//rate limiter
  func AuthLimiter(ctx *fiber.Ctx) error {
	limiter.New(limiter.Config{
		Max: 60, //Maximum Request
		Expiration: 40 * time.Second, //Maximum Dalam waktu 40 detik 
		LimitReached: func (c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).SendStatus(429)
		},
	})
 return ctx.Next()
}







