package routes

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
	middlewares "shagya-tech-payment/api/middleware"
	"time"
)

var (
	Router *fiber.App
)

func Route() {
	Router.Static("/demo", "./public/views")
	Router.Use(middlewares.CORSMiddleware())
	api := Router.Group("/api-payment")
	v1 := api.Group("/v.1")

	master := v1.Group("/master")
	master.Use(middlewares.RateLimitMiddleware(5, 10*time.Second))
	{
		master.Get("/product", ProductController.Product)
	}
	v1.Get("/storage/product/:filename", func(c *fiber.Ctx) error {
		pwd, _ := os.Getwd()
		filename := c.Params("filename")
		return c.SendFile(fmt.Sprintf("%s/public/product/%s.png", pwd, filename))
	})
}
