package controller

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"shagya-tech-payment/internal/service"
	"shagya-tech-payment/pkg"
)

type ProductController struct {
	DB             *mongo.Database
	Client         *mongo.Client
	ProductService service.ProductService
}

func HandlerProductController(db *mongo.Database, client *mongo.Client) *ProductController {
	if db == nil {
		log.Println("Database [HandlerController] connection is nil")
	}

	return &ProductController{
		DB:     db,
		Client: client,
		ProductService: service.ProductService{
			DB:     db,
			Client: client,
		},
	}
}

func (ct *ProductController) Product(c *fiber.Ctx) error {
	responseInitial := pkg.InitialResponse{Ctx: c}

	result, err := ct.ProductService.Product(c.Query("category"), c.Query("brand"), c.Query("type"), c.Query("code"))
	if err != nil {
		return responseInitial.Respose(fiber.StatusBadRequest, err.Error(), true, nil)
	}
	return responseInitial.Respose(http.StatusOK, "Ok", false, result)
}
