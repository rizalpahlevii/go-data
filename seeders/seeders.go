package seeders

import (
	"fmt"
)

func RunSeeders() {
	fmt.Println("Running seeders...")
	//start := time.Now()
	//
	//locationsSeeder()
	//geolocationsSeeder()
	//customersSeeder()
	//sellersSeeder()
	//productsSeeder()
	//ordersSeeder()
	//orderItemsSeeder()
	//orderPaymentsSeeder()
	orderReviewsSeeder()
	//fmt.Println("Seeders finished. time elapsed: " + time.Since(start).String())

}
