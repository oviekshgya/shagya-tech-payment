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

		var jsonResponse models.MasterDataProduct
		err = json.Unmarshal(fileData, &jsonResponse)
		if err != nil {
			log.Println("Error unmarshaling JSON:", err)
		}

		model := models.CronsImpl{
			DB:                          s.DB,
			MasterDataProductCollection: jsonResponse,
		}

		if created := model.Create(); created != nil {
			return created, nil
		}

		return jsonResponse, nil
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}
