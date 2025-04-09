package routes

import (
	"github.com/gofiber/fiber/v2"
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
}
