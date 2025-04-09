package controller

import (
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"shagya-tech-payment/internal/service"
)

type CronsController struct {
	DB              *mongo.Database
	Client          *mongo.Client
	CronsController service.CronsService
}

func HandlerCronsController(db *mongo.Database, client *mongo.Client) *CronsController {
	if db == nil {
		log.Println("Database [HandlerController] connection is nil")
	}

	return &CronsController{
		DB:     db,
		Client: client,
		CronsController: service.CronsService{
			DB:     db,
			Client: client,
		},
	}
}

func (ct *CronsController) GetDataJson() bool {
	_, err := ct.CronsController.GetDataJson()
	if err != nil {
		log.Println("CronsController.GetDataJson ERROR:", err)
		return false
	}
	return true
}
