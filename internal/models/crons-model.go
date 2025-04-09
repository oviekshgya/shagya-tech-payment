package models

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
)

const MASTERDATAPRODUCT = "master-data-product"

type CronsImpl struct {
	DB                          *mongo.Database
	MasterDataProductCollection MasterDataProduct
}
type MasterDataProduct struct {
	Name     string `bson:"name" json:"name"`
	ImageURL string `bson:"imageURL" json:"image_url"`
}

func (MasterDataProduct) TableName() string {
	return MASTERDATAPRODUCT
}

func (c *CronsImpl) Create() error {
	_, err := c.DB.Collection(MASTERDATAPRODUCT).InsertOne(context.Background(), c.MasterDataProductCollection)
	if err != nil {
		return fmt.Errorf("failed to insert user: %v", err)
	}
	return nil
}
