package service

import (
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"io/ioutil"
	"log"
	"os"
	"shagya-tech-payment/internal/models"
	"shagya-tech-payment/pkg"
	"strings"
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

	pwd, _ := os.Getwd()
	fileData, err := ioutil.ReadFile(fmt.Sprintf("%s/public/storage/json/tokogame-data.json", pwd))
	if err != nil {
		log.Println("Error reading file:", err)
	}

	var jsonResponse []models.MasterDataProduct
	err = json.Unmarshal(fileData, &jsonResponse)
	if err != nil {
		log.Println("Error unmarshaling JSON:", err)
	}

	var response []*pkg.ResponseProduct
	if len(result.Data) > 0 {
		for _, d := range result.Data {
			var imageURL string
			for _, sz := range jsonResponse {
				matchCount := 0

				words := strings.Fields(sz.Name)

				for _, word := range words {
					if strings.Contains(strings.ToLower(d.Brand), strings.ToLower(word)) {
						matchCount++
					}
				}
				if matchCount >= 2 {
					imageURL = sz.ImageURL
					break
				}
			}

			response = append(response, &pkg.ResponseProduct{
				ProductName:   d.ProductName,
				Category:      d.Category,
				Image:         imageURL,
				Price:         d.Price,
				Stock:         d.Stock,
				Desc:          d.Desc,
				MiddlewareFee: 0,
			})
		}
	}
	return response, nil
}
