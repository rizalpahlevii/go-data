package migrations

import (
	"context"
	"fmt"
	"main/database"
)

func runOrder() {
	orderDown()
	orderUp()
}

func orderUp() {
	db := database.GetConnection()
	defer db.Close()

	ctx := context.Background()

	// Create table
	createQuery := `CREATE TABLE orders(order_id varchar(50) NOT NULL, customer_id varchar(50) NOT NULL, order_status varchar(15) NOT NULL, order_purchase_timestamp TIMESTAMP NOT NULL, order_approved_at TIMESTAMP, order_delivered_carrier_date TIMESTAMP, order_delivered_customer_date TIMESTAMP, order_estimated_delivery_date TIMESTAMP NOT NULL, PRIMARY KEY (order_id))`

	_, err := db.ExecContext(ctx, createQuery)
	if err != nil {
		panic(err)
	}
	fmt.Println("Table order created successfully")
}

func orderDown() {
	db := database.GetConnection()
	defer db.Close()

	ctx := context.Background()

	// Drop table
	dropQuery := `DROP TABLE IF EXISTS orders`

	_, err := db.ExecContext(ctx, dropQuery)
	if err != nil {
		panic(err)
	}
	fmt.Println("Table orders dropped successfully")
}
