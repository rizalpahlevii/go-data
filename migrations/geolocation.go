package migrations

import (
	"context"
	"fmt"
	"main/database"
)

func runGeolocation() {
	geolocationDown()
	geolocationUp()
}

//"geolocation_zip_code_prefix","geolocation_lat","geolocation_lng","geolocation_city","geolocation_state"
func geolocationDown() {
	db := database.GetConnection()
	defer db.Close()

	ctx := context.Background()

	// Drop table
	dropQuery := "DROP TABLE IF EXISTS geolocation"
	_, err := db.ExecContext(ctx, dropQuery)
	if err != nil {
		panic(err)
	}
	fmt.Println("Table geolocation dropped successfully")
}

func geolocationUp() {
	db := database.GetConnection()
	defer db.Close()

	ctx := context.Background()

	// Create table
	script := "CREATE TABLE geolocation (geolocation_zip_code_prefix VARCHAR(50), geolocation_lat VARCHAR(50), geolocation_lng VARCHAR(50) NULL, geolocation_city VARCHAR(50) NULL, geolocation_state VARCHAR(50) NULL)"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("Table geolocation created successfully")
}
