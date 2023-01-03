package seeders

import (
	"context"
	"encoding/csv"
	"fmt"
	"main/database"
	"os"
	"time"
)

func customersSeeder() {
	fmt.Println("Importing customers...")
	start := time.Now()

	csvFile, err := os.Open("datasets/olist_customers_dataset.csv")
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
		tx, err := db.Begin()
		if err != nil {
			panic(err)
		}

		stmt, err := tx.Prepare("INSERT INTO customer ( customer_id, customer_unique_id, customer_zip_code_prefix, customer_city, customer_state) VALUES (?, ?, ?, ?, ?)")
		if err != nil {
			panic(err)
		}
		defer stmt.Close()

		// iterate over the records start from 1 to skip the header

		for i := 1; i < len(records); i++ {
			//	"customer_id","customer_unique_id","customer_zip_code_prefix","customer_city","customer_state"
			record := records[i]

			_, err = stmt.ExecContext(ctx, record[0], record[1], record[2], record[3], record[4])
			if err != nil {
				panic(err)
			}
		}

		if err = tx.Commit(); err != nil {
			panic(err)
		}
		done <- true
	}()

	// wait for the insert to complete
	<-done
	fmt.Printf("Imported %d records in %s", len(records)-1, time.Since(start))

}
