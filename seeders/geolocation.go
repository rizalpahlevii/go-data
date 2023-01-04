package seeders

import (
	"context"
	"encoding/csv"
	"fmt"
	"main/configurations"
	"main/database"
	"os"
	"time"
)

func geolocationsSeeder() {
	fmt.Println("Importing geolocations...")
	start := time.Now()

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

	// open connection to the database
	db := database.GetConnection()
	defer db.Close()
	ctx := context.Background()

	// create a channel to communicate completion
	done := make(chan bool)

	// start the insert in a separate goroutine
	go func() {

		// Begin a transaction
		tx, err := db.Begin()
		if err != nil {
			panic(err)
		}

		// prepare the insert statement
		stmt, err := tx.Prepare("INSERT INTO " + configurations.Configuration().GeolocationTableName + " ( geolocation_zip_code_prefix, geolocation_lat, geolocation_lng, geolocation_city, geolocation_state) VALUES (?, ?, ?, ?, ?)")
		if err != nil {
			panic(err)
		}
		defer stmt.Close()

		// iterate over the records start from 1 to skip the header
		for i := 1; i < len(records); i++ {
			//	geolocation_zip_code_prefix","geolocation_lat","geolocation_lng","geolocation_city","geolocation_state"
			record := records[i]

			_, err = stmt.ExecContext(ctx, record[0], record[1], record[2], record[3], record[4])
			if err != nil {
				panic(err)
			}
		}

		if err := tx.Commit(); err != nil {
			panic(err)
		}

		done <- true
	}()

	// wait for the insert to complete
	<-done

	fmt.Printf("Imported %d records of customer in %s", len(records)-1, time.Since(start))
}
