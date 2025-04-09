package service

import (
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"shagya-tech-payment/internal/models"
	"shagya-tech-payment/pkg"
	"strings"
)

type CronsService struct {
	DB     *mongo.Database
	Client *mongo.Client
}

func (s *CronsService) GetDataJson() (interface{}, error) {

	result, err := pkg.WithTransactionMongo(s.Client, func(sessCtx mongo.SessionContext) (interface{}, error) {

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

		if len(jsonResponse) > 0 {
			for i, product := range jsonResponse {
				if strings.Contains(product.ImageURL, "https") {
					resp, err := http.Get(product.ImageURL)
					if err != nil {
						panic(err)
					}
					defer resp.Body.Close()

					imageName := pkg.GenerateMD5(product.ImageURL)

					out, err := os.Create(fmt.Sprintf("%s/public/product/%s.png", pwd, imageName))
					if err != nil {
						panic(err)
					}
					defer out.Close()

					_, err = io.Copy(out, resp.Body)
					if err != nil {
						panic(err)
					}

					jsonResponse[i].ImageURL = fmt.Sprintf("https://shagya-tech.my.id/api-payment/v.1/storage/product/%s", imageName)

				}
			}
		}

		model := models.CronsImpl{
			DB:                          s.DB,
			MasterDataProductCollection: jsonResponse,
		}

		data, errGet := model.GetAll()
		if errGet == nil && len(data) <= 0 {
			if created := model.Create(jsonResponse); created != nil {
				return created, nil
			}
		}

		return jsonResponse, nil
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}
