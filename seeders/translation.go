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

func translationSeeder() {
	fmt.Println("Importing translations...")

	start := time.Now()

	// open the file
	csvFile, err := os.Open("datasets/product_category_name_translation.csv")
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
		stmt, err := tx.Prepare("INSERT INTO " + configurations.Configuration().TranslationTableName + " ( product_category_name, product_category_name_english) VALUES (?, ?)")
		if err != nil {
			panic(err)
		}
		defer stmt.Close()

		// iterate over the records start from 1 to skip the header
		for i := 1; i < len(records); i++ {
			//	"product_category_name","product_category_name_english"
			record := records[i]

			_, err = stmt.ExecContext(ctx, record[0], record[1])
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

	fmt.Printf("Imported %d records of product translation in %s", len(records)-1, time.Since(start))

}
