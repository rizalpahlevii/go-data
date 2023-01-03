package seeders

import (
	"context"
	"encoding/csv"
	"fmt"
	"main/database"
	"os"
	"time"
)

func orderPaymentsSeeder() {
	fmt.Println("Importing order payments...")

	start := time.Now()

	csvFile, err := os.Open("datasets/olist_order_payments_dataset.csv")
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

	go func() {
		tx, err := db.Begin()
		if err != nil {
			panic(err)
		}

		stmt, err := tx.Prepare("INSERT INTO order_payment ( order_id, payment_sequential, payment_type, payment_installments, payment_value) VALUES (?, ?, ?, ?, ?)")
		if err != nil {
			panic(err)
		}
		defer stmt.Close()

		// iterate over the records start from 1 to skip the header
		for i := 1; i < len(records); i++ {
			// "order_id","payment_sequential","payment_type","payment_installments","payment_value"
			record := records[i]
			orderId := record[0]
			paymentSequential := record[1]
			paymentType := record[2]
			paymentInstallments := record[3]
			paymentValue := record[4]

			_, err = stmt.ExecContext(ctx, orderId, paymentSequential, paymentType, paymentInstallments, paymentValue)
			if err != nil {
				panic(err)
			}
		}

		err = tx.Commit()
		if err != nil {
			panic(err)
		}

		done <- true
	}()

	<-done
	fmt.Printf("Imported %d records of order payment in %s", len(records)-1, time.Since(start))
}
