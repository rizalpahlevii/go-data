package configurations

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

var shared *_Configuration

type _Configuration struct {
	LocationTableName     string `json:"location_table_name"`
	GeolocationTableName  string `json:"geolocation_table_name"`
	OrderTableName        string `json:"order_table_name"`
	OrderItemTableName    string `json:"order_item_table_name"`
	OrderPaymentTableName string `json:"order_payment_table_name"`
	OrderReviewTableName  string `json:"order_review_table_name"`
	SellerTableName       string `json:"seller_table_name"`
	CustomerTableName     string `json:"customer_table_name"`
	ProductTableName      string `json:"product_table_name"`
	TranslationTableName  string `json:"translation_table_name"`
}

func init() {

	if shared != nil {
		return
	}

	basePath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// Load the configuration file
	bts, err := ioutil.ReadFile(filepath.Join(basePath, "configurations", "config.json"))

	if err != nil {
		panic(err)
	}

	shared = new(_Configuration)
	err = json.Unmarshal(bts, shared)
	if err != nil {
		panic(err)
		return
	}
}

func Configuration() _Configuration {
	return *shared
}
