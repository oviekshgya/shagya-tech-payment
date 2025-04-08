package service

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	DB     *mongo.Database
	Client *mongo.Client
}
