package service

import (
	"go.mongodb.org/mongo-driver/mongo"
	"shagya-tech-payment/pkg"
)

type ProductService struct {
	DB     *mongo.Database
	Client *mongo.Client
}

func (s *ProductService) Product(category string, brand string, types string, code string) (*pkg.ResponseCekHarga, error) {

	result, err := pkg.CekHargaRest(pkg.RequestCekHarga{
		Category: category,
		Brand:    brand,
		Username: "zicubaDRnGAD",
		Type:     types,
		Cmd:      "prepaid",
		Sign:     pkg.GenerateMD5("zicubaDRnGADdev-5560e150-138a-11f0-9b7a-0337a9521aacpricelist"), //Signature dengan formula md5(username + apiKey + "pricelist")
		Code:     code,
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}
