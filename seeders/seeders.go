package seeders

import (
	"fmt"
	"time"
)

func RunSeeders() {
	fmt.Println("Running seeders...")
	start := time.Now()

	locationsSeeder()
	geolocationsSeeder()
	customersSeeder()
	sellersSeeder()
	productsSeeder()
	ordersSeeder()
	orderItemsSeeder()
	orderPaymentsSeeder()
	orderReviewsSeeder()
	translationSeeder()

	fmt.Println("Seeders finished. time elapsed: " + time.Since(start).String())

}
