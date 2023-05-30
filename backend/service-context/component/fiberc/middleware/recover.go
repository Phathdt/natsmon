package middleware

import (
	"github.com/gofiber/fiber/v2"
	sctx "natsmon/service-context"
)

func Recover(sc sctx.ServiceContext) fiber.Handler {
	// Return new handler
	return func(ctx *fiber.Ctx) error {
		defer func() {
			if err := recover(); err != nil {
				if err = ctx.Status(500).JSON(&fiber.Map{
					"errors": err,
				}); err != nil {
					return
				}
			}
		}()

		// Return err if existed, else move to next handler
		return ctx.Next()
	}
}
