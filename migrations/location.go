package migrations

import (
	"context"
	"fmt"
	"main/database"
)

func runLocation() {
	locationDown()
	locationUp()
}

func locationDown() {
	db := database.GetConnection()
	defer db.Close()
	ctx := context.Background()
	// Drop table if exists
	dropQuery := "DROP TABLE IF EXISTS location"
	_, err := db.ExecContext(ctx, dropQuery)
	if err != nil {
		panic(err)
	}
	fmt.Println("Table location dropped successfully")
}

// Create table location base on csv file in datasets/locations.csv
//state,state_code,city,zip_code_start,zip_code_end,state_lat,state_lng
func locationUp() {
	db := database.GetConnection()
	defer db.Close()
	ctx := context.Background()

	// Create table location
	script := "CREATE TABLE location (state VARCHAR(255), state_code VARCHAR(255), city VARCHAR(255), zip_code_start VARCHAR(255), zip_code_end VARCHAR(255), state_lat VARCHAR(255), state_lng VARCHAR(255))"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("Table location created successfully")
}
