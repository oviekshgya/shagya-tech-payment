package routes

import (
	"shagya-tech-payment/api/controller"
	"shagya-tech-payment/db"
)

var (
	UserController *controller.UserController
)

func InitialRoute() {
	UserController = controller.HandlerController(db.DBMongo, db.Client)
}
