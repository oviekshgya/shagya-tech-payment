package controller

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"shagya-tech-payment/internal/service"
	"shagya-tech-payment/pkg"
)

type UserController struct {
	DB          *mongo.Database
	Client      *mongo.Client
	UserService service.UserService
}

func HandlerController(db *mongo.Database, client *mongo.Client) *UserController {
	if db == nil {
		log.Println("Database [HandlerController] connection is nil")
	}

	return &UserController{
		DB:     db,
		Client: client,
		UserService: service.UserService{
			DB:     db,
			Client: client,
		},
	}
}

func (controller *UserController) Welcome(c *fiber.Ctx) error {
	responseInitial := pkg.InitialResponse{Ctx: c}
	return responseInitial.Respose(http.StatusOK, "Welcome", false, map[string]interface{}{
		"message": "Welcome Shagya-Tech Payment" + pkg.GoogleOAuthConfig.ClientID,
	})
}
