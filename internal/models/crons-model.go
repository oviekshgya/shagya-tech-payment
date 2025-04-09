package models

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const MASTERDATAPRODUCT = "master-data-product"

type CronsImpl struct {
	DB                          *mongo.Database
	MasterDataProductCollection []MasterDataProduct
}
type MasterDataProduct struct {
	Name     string `bson:"name" json:"name"`
	ImageURL string `bson:"imageURL" json:"image_url"`
}

func (MasterDataProduct) TableName() string {
	return MASTERDATAPRODUCT
}

func (c *CronsImpl) Create(data []MasterDataProduct) error {

	docs := make([]interface{}, len(data))
	if len(data) > 0 {
		for i, v := range data {
			docs[i] = v
		}
	}
	_, err := c.DB.Collection(MASTERDATAPRODUCT).InsertMany(context.Background(), docs)
	if err != nil {
		return fmt.Errorf("failed to insert user: %v", err)
	}
	return nil
}

func (c *CronsImpl) GetAll() ([]MasterDataProduct, error) {
	cursor, err := c.DB.Collection(MASTERDATAPRODUCT).Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %v", err)
	}
	var results []MasterDataProduct
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, fmt.Errorf("failed to find user all: %v", err)
	}

	return results, nil
}
