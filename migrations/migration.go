package migrations

func RunMigrations() {
	runLocation()
	runGeolocation()
	runOrder()
	runOrderItem()
	runOrderPayment()
	runOrderReview()
	runSeller()
	runProduct()
	runTranslation()
	runCustomer()
}
