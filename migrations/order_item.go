package migrations

import (
	"context"
	"fmt"
	"main/configurations"
	"main/database"
)

//"order_id","order_item_id","product_id","seller_id","shipping_limit_date","price","freight_value"
func runOrderItem() {
	orderItemDown()
	orderItemUp()
}

func orderItemDown() {
	db := database.GetConnection()
	defer db.Close()

	ctx := context.Background()

	dropQuery := "DROP TABLE IF EXISTS " + configurations.Configuration().OrderItemTableName
	_, err := db.ExecContext(ctx, dropQuery)
	if err != nil {
		panic(err)
	}
	fmt.Println("Table order items dropped successfully")
}

func orderItemUp() {
	db := database.GetConnection()
	defer db.Close()

	ctx := context.Background()

	createQuery := `CREATE TABLE ` + configurations.Configuration().OrderItemTableName + ` (
		order_id varchar(50) NOT NULL,
		order_item_id varchar(50) NOT NULL,
		product_id varchar(50) NOT NULL,
		seller_id varchar(50) NOT NULL,
		shipping_limit_date TIMESTAMP NOT NULL,
		price FLOAT NOT NULL,
		freight_value FLOAT NOT NULL)`

	_, err := db.ExecContext(ctx, createQuery)
	if err != nil {
		panic(err)
	}
	fmt.Println("Table order items created successfully")
}
