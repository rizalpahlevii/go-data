package seeders

import (
	"context"
	"encoding/csv"
	"fmt"
	"main/database"
	"os"
)

func ImportGeolocations() {
	fmt.Println("Importing geolocations...")

	// open connection to the database
	db := database.GetConnection()
	defer db.Close()

	ctx := context.Background()

	// open the file
	csvFile, err := os.Open("datasets/olist_geolocation_dataset.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	// parse the file
	r := csv.NewReader(csvFile)
	records, err := r.ReadAll()
	if err != nil {
		panic(err)
	}

	// iterate over the records start from 1 to skip the header
	for i := 1; i < len(records); i++ {
		//	geolocation_zip_code_prefix","geolocation_lat","geolocation_lng","geolocation_city","geolocation_state"
		record := records[i]
		query := "INSERT INTO geolocations (geolocation_zip_code_prefix,geolocation_lat,geolocation_lng,geolocation_city,geolocation_state) VALUES (?, ?, ?, ?, ?)"
		_, err = db.ExecContext(ctx, query, record[0], record[1], record[2], record[3], record[4])
		if err != nil {
			panic(err)
		}

	}

	fmt.Println("Geolocations imported successfully")
}
