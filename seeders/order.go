package seeders

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"main/database"
	"os"
	"time"
)

func ordersSeeder() {
	fmt.Println("Importing orders...")

	start := time.Now()

	csvFile, err := os.Open("datasets/olist_orders_dataset.csv")
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
	//ctx := context.Background()

	// create a channel to communicate completion
	done := make(chan bool)

	go func() {
		tx, err := db.Begin()
		if err != nil {
			panic(err)
		}

		stmt, err := tx.Prepare("INSERT INTO orders ( order_id, customer_id, order_status, order_purchase_timestamp, order_approved_at, order_delivered_carrier_date, order_delivered_customer_date, order_estimated_delivery_date) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
		if err != nil {
			panic(err)
		}
		defer stmt.Close()

		// iterate over the records start from 1 to skip the header

		for i := 1; i < len(records); i++ {
			//	"order_id","customer_id","order_status","order_purchase_timestamp","order_approved_at","order_delivered_carrier_date","order_delivered_customer_date","order_estimated_delivery_date"
			record := records[i]
			orderId := record[0]
			customerId := record[1]
			orderStatus := record[2]
			orderPurchaseTimestamp := record[3]
			orderApprovedAt := record[4]
			orderDeliveredCarrierDate := record[5]
			orderDeliveredCustomerDate := record[6]
			orderEstimatedDeliveryDate := record[7]

			var orderApprovedAtNull sql.NullTime
			if orderApprovedAt != "" {
				t, err := time.Parse("2006-01-02 15:04:05", orderApprovedAt)
				if err != nil {
					panic(err)
				}
				orderApprovedAtNull.Time = t
				orderApprovedAtNull.Valid = true
			}

			var orderDeliveredCarrierDateNull sql.NullTime
			if orderDeliveredCarrierDate != "" {
				t, err := time.Parse("2006-01-02 15:04:05", orderDeliveredCarrierDate)
				if err != nil {
					panic(err)
				}
				orderDeliveredCarrierDateNull.Time = t
				orderDeliveredCarrierDateNull.Valid = true
			}

			var orderDeliveredCustomerDateNull sql.NullTime
			if orderDeliveredCustomerDate != "" {
				t, err := time.Parse("2006-01-02 15:04:05", orderDeliveredCustomerDate)
				if err != nil {
					panic(err)
				}
				orderDeliveredCustomerDateNull.Time = t
				orderDeliveredCustomerDateNull.Valid = true
			}

			_, err = stmt.Exec(orderId, customerId, orderStatus, orderPurchaseTimestamp, orderApprovedAtNull, orderDeliveredCarrierDateNull, orderDeliveredCustomerDateNull, orderEstimatedDeliveryDate)
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

	fmt.Printf("Imported %d records of order in %s", len(records)-1, time.Since(start))

}
