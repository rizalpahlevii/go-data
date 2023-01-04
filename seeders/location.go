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

func locationsSeeder() {
	fmt.Println("Importing locations...")
	start := time.Now()

	// open connection to the database
	db := database.GetConnection()
	defer db.Close()

	ctx := context.Background()

	csvFile, err := os.Open("datasets/locations.csv")
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
		record := records[i]

		// insert into the database
		query := "INSERT INTO " + configurations.Configuration().LocationTableName + " (state, state_code, city, zip_code_start, zip_code_end, state_lat, state_lng) VALUES (?, ?, ?, ?, ?, ?, ?)"
		_, err = db.ExecContext(ctx, query, record[0], record[1], record[2], record[3], record[4], record[5], record[6])
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Imported records:", len(records)-1)
	fmt.Println("Locations imported successfully with time elapsed: ", time.Since(start))

}
