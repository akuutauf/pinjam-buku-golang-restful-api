package middleware

import "github.com/gofiber/fiber/v2"

func AuthMiddleware(ctx *fiber.Ctx) error {
	return ctx.Next()
}
