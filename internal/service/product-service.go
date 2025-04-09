package service

import (
	"go.mongodb.org/mongo-driver/mongo"
	"shagya-tech-payment/pkg"
)

type ProductService struct {
	DB     *mongo.Database
	Client *mongo.Client
}

func (s *ProductService) Product(category string, brand string, types string, code string) ([]*pkg.ResponseProduct, error) {

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

	var response []*pkg.ResponseProduct
	if len(result.Data) > 0 {
		for _, d := range result.Data {
			response = append(response, &pkg.ResponseProduct{
				ProductName:   d.ProductName,
				Category:      d.Category,
				Image:         "",
				Price:         d.Price,
				Stock:         d.Stock,
				Desc:          d.Desc,
				MiddlewareFee: 0,
			})
		}
	}
	return response, nil
}
