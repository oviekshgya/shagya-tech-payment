package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func debugCurlRequest(req *http.Request, body []byte) string {
	var curlCommand strings.Builder
	curlCommand.WriteString("curl -X ")
	curlCommand.WriteString(req.Method)
	curlCommand.WriteString(fmt.Sprintf(" \"%s\"", req.URL.String()))

	for key, values := range req.Header {
		for _, value := range values {
			curlCommand.WriteString(fmt.Sprintf(" -H \"%s: %s\"", key, value))
		}
	}

	if len(body) > 0 {
		curlCommand.WriteString(fmt.Sprintf(" -d '%s'", body))
	}

	fmt.Println("Curl Debug:", curlCommand.String())
	return curlCommand.String()
}

type ResponseCekHarga struct {
	Data []struct {
		ProductName         string `json:"product_name"`
		Category            string `json:"category"`
		Brand               string `json:"brand"`
		Type                string `json:"type"`
		SellerName          string `json:"seller_name"`
		Price               int    `json:"price"`
		BuyerSkuCode        string `json:"buyer_sku_code"`
		BuyerProductStatus  bool   `json:"buyer_product_status"`
		SellerProductStatus bool   `json:"seller_product_status"`
		UnlimitedStock      bool   `json:"unlimited_stock"`
		Stock               int    `json:"stock"`
		Multi               bool   `json:"multi"`
		StartCutOff         string `json:"start_cut_off"`
		EndCutOff           string `json:"end_cut_off"`
		Desc                string `json:"desc"`
	} `json:"data"`
}
type RequestCekHarga struct {
	Cmd      string `json:"cmd"`
	Username string `json:"username"`
	Sign     string `json:"sign"`
	Category string
	Brand    string `json:"brand"`
	Type     string `json:"type"`
	Code     string `json:"code"`
}

func CekHargaRest(req RequestCekHarga) (*ResponseCekHarga, error) {
	client := &http.Client{}
	var data map[string]interface{}

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}
	var jsonStr = []byte(string(requestBody))

	hit, err2 := http.NewRequest("POST", "https://api.digiflazz.com/v1/price-list", bytes.NewBuffer(jsonStr))
	if err2 != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	hit.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(hit)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer func(Body io.ReadCloser) {
		errBody := Body.Close()
		if errBody != nil {
			fmt.Println("errBody", errBody)
		}
	}(resp.Body)

	errJson := json.NewDecoder(resp.Body).Decode(&data)
	if errJson != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", errJson)
	}

	jsonData, _ := json.Marshal(data)

	var result ResponseCekHarga
	json.Unmarshal(jsonData, &result)

	debugCurlRequest(hit, jsonStr)

	return &result, nil
}
