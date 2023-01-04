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

func productsSeeder() {
	fmt.Println("Importing products...")
	start := time.Now()

	csvFile, err := os.Open("datasets/olist_products_dataset.csv")
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

		stmt, err := tx.Prepare("INSERT INTO " + configurations.Configuration().ProductTableName + " ( product_id, product_category_name, product_name_length, product_description_length, product_photos_qty, product_weight_g, product_length_cm, product_height_cm, product_width_cm) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)")
		if err != nil {
			panic(err)
		}
		defer stmt.Close()

		// iterate over the records start from 1 to skip the header

		for i := 1; i < len(records); i++ {
			//	"product_id","product_category_name","product_name_length","product_description_length","product_photos_qty","product_weight_g","product_length_cm","product_height_cm","product_width_cm"
			// check if the record is empty or not, if it is empty change it to NULL
			record := records[i]
			for j := 0; j < len(record); j++ {
				if record[j] == "" {
					record[j] = "0"
				}
			}

			_, err = stmt.ExecContext(ctx, record[0], record[1], record[2], record[3], record[4], record[5], record[6], record[7], record[8])
			if err != nil {
				panic(err)
			}
		}
		tx.Commit()
		done <- true
	}()

	// wait for the insert to complete
	<-done

	fmt.Printf("Imported %d records of customer in %s", len(records)-1, time.Since(start))

}
