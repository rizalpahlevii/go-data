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

func sellersSeeder() {
	fmt.Println("Importing sellers...")

	start := time.Now()
	csvFile, err := os.Open("datasets/olist_sellers_dataset.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	// parse the file
	records, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		panic(err)
	}

	// open connection to the database
	db := database.GetConnection()
	defer db.Close()
	ctx := context.Background()

	// create a channel to communicate completion
	done := make(chan bool)

	//	start the insert in a separate goroutine
	go func() {

		// Begin a transaction
		tx, err := db.Begin()
		if err != nil {
			panic(err)
		}

		// prepare the insert statement
		stmt, err := tx.Prepare("INSERT INTO " + configurations.Configuration().SellerTableName + " ( seller_id, seller_zip_code_prefix, seller_city, seller_state) VALUES (?, ?, ?, ?)")
		if err != nil {
			panic(err)
		}
		defer stmt.Close()

		// iterate over the records start from 1 to skip the header
		for i := 1; i < len(records); i++ {
			//	"seller_id","seller_zip_code_prefix","seller_city","seller_state"
			record := records[i]

			_, err = stmt.ExecContext(ctx, record[0], record[1], record[2], record[3])
			if err != nil {
				panic(err)
			}
		}

		// commit the transaction
		err = tx.Commit()
		if err != nil {
			panic(err)
		}

		// signal completion
		done <- true
	}()

	// wait for completion
	<-done

	fmt.Printf("Imported %d sellers in %s", len(records)-1, time.Since(start))
}
