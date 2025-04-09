package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func CORSMiddleware() fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins: "https://shagyaapi-production.up.railway.app", //Deployment
		//AllowOrigins:     "http://127.0.0.1", //Deployment
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, X-API-KEY, X-SIGNATURE, X-TIMESTAMP",
		AllowCredentials: true,
	})
}

func APIKeyMiddleware(validAPIKey string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		apiKey := c.Get("X-API-KEY")

		if apiKey == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "API Key is missing",
			})
		}

		if apiKey != validAPIKey {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Invalid API Key",
			})
		}

		return c.Next()
	}
}
