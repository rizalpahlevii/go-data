package main

import (
	"main/configurations"
	"main/database"
)

func getRevenues() {
	locationTable := configurations.Configuration().CustomerTableName
	geolocationTable := configurations.Configuration().GeolocationTableName
	orderTable := configurations.Configuration().OrderTableName
	orderItemTable := configurations.Configuration().OrderItemTableName
	productTable := configurations.Configuration().ProductTableName
	translationTable := configurations.Configuration().TranslationTableName
	orderPaymentTable := configurations.Configuration().OrderPaymentTableName
	orderReviewTable := configurations.Configuration().OrderReviewTableName
	sellerTable := configurations.Configuration().SellerTableName
	customerTable := configurations.Configuration().CustomerTableName

	db := database.GetConnection()
	defer db.Close()

	//	Displays the revenue (profit) earned every day based on sales every day by looking at the price paid from datasets

}
