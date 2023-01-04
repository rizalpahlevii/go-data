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

func orderItemsSeeder() {

	fmt.Println("Importing order items...")
	start := time.Now()

	csvFile, err := os.Open("datasets/olist_order_items_dataset.csv")
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

		stmt, err := tx.Prepare("INSERT INTO " + configurations.Configuration().OrderItemTableName + " ( order_id, order_item_id, product_id, seller_id, shipping_limit_date, price, freight_value) VALUES (?, ?, ?, ?, ?, ?, ?)")
		if err != nil {
			panic(err)
		}
		defer stmt.Close()

		// iterate over the records start from 1 to skip the header
		for i := 1; i < len(records); i++ {
			// "order_id","order_item_id","product_id","seller_id","shipping_limit_date","price","freight_value"
			record := records[i]
			orderId := record[0]
			orderItemId := record[1]
			productId := record[2]
			sellerId := record[3]
			shippingLimitDate := record[4]
			price := record[5]
			freightValue := record[6]

			_, err = stmt.ExecContext(ctx, orderId, orderItemId, productId, sellerId, shippingLimitDate, price, freightValue)
			if err != nil {
				panic(err)
			}
		}

		if err := tx.Commit(); err != nil {
			panic(err)
		}

		done <- true
	}()

	<-done

	fmt.Printf("Imported %d records of order item in %s", len(records)-1, time.Since(start))
}
