package migrations

import (
	"context"
	"fmt"
	"main/database"
)

//"order_id","payment_sequential","payment_type","payment_installments","payment_value"

func runOrderPayment() {
	orderPaymentDown()
	orderPaymentUp()
}

func orderPaymentUp() {
	db := database.GetConnection()
	defer db.Close()

	ctx := context.Background()
	query := "CREATE TABLE order_payment (order_id varchar(50), payment_sequential int, payment_type varchar(255), payment_installments int, payment_value float);"
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}
	fmt.Println("Table order payment created")
}

func orderPaymentDown() {
	db := database.GetConnection()
	defer db.Close()

	ctx := context.Background()
	query := "DROP TABLE IF EXISTS order_payment;"
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}
	fmt.Println("Table order payment deleted")
}
