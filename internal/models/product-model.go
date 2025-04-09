package models

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const PRODUCT = "product"

type ProductImpl struct {
	DB                          *mongo.Database
	MasterDataProductCollection []MasterDataProduct
}

type Product struct {
	ProductName         string `bson:"productName" json:"product_name"`
	Category            string `bson:"category" json:"category"`
	Brand               string `bson:"brand" json:"brand"`
	Type                string `bson:"type" json:"type"`
	SellerName          string `bson:"sellerName" json:"seller_name"`
	Price               int    `bson:"price" json:"price"`
	BuyerSkuCode        string `bson:"buyerSkuCode" json:"buyer_sku_code"`
	BuyerProductStatus  bool   `bson:"buyerProductStatus" json:"buyer_product_status"`
	SellerProductStatus bool   `bson:"sellerProductStatus" json:"seller_product_status"`
	UnlimitedStock      bool   `bson:"unlimitedStock" json:"unlimited_stock"`
	Stock               int    `bson:"stock" json:"stock"`
	Multi               bool   `bson:"multi" json:"multi"`
	StartCutOff         string `bson:"startCutOff" json:"start_cut_off"`
	EndCutOff           string `bson:"endCutOff" json:"end_cut_off"`
	Desc                string `bson:"desc" json:"desc"`
}

func (Product) TableName() string {
	return PRODUCT
}

func (m *ProductImpl) GetProductByName(productName string) (*Product, error) {
	var data Product
	err := m.DB.Collection(PRODUCT).FindOne(context.Background(), bson.M{"productName": productName}).Decode(&data)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("productName not found")
		}
		return nil, fmt.Errorf("failed to get productName: %v", err)
	}
	return &data, nil
}
