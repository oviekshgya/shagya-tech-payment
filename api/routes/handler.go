package routes

import (
	"shagya-tech-payment/api/controller"
	"shagya-tech-payment/db"
)

var (
	UserController  *controller.UserController
	CronsController *controller.CronsController
)

func InitialRoute() {
	UserController = controller.HandlerController(db.DBMongo, db.Client)
	CronsController = controller.HandlerCronsController(db.DBMongo, db.Client)
}
