package migrations

import "fmt"

func RunMigrations() {
	fmt.Println("Running migrations...")
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
	fmt.Println("Migrations run successfully")
}
