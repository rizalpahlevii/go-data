package migrations

import (
	"context"
	"fmt"
	"main/database"
)

func runSeller() {
	sellerDown()
	sellerUp()
}

//"seller_id","seller_zip_code_prefix","seller_city","seller_state"
func sellerUp() {
	db := database.GetConnection()
	defer db.Close()
	ctx := context.Background()

	// Your code here
	query := `CREATE TABLE seller(
		seller_id varchar(50) NOT NULL,
		seller_zip_code_prefix VARCHAR(7) NOT NULL,
		seller_city VARCHAR(255),
		seller_state VARCHAR(255),
		PRIMARY KEY (seller_id)
	)`
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}
	fmt.Println("Table seller created successfully")
}

func sellerDown() {
	db := database.GetConnection()
	defer db.Close()
	ctx := context.Background()

	// Your code here
	query := `DROP TABLE IF EXISTS seller`
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}
	fmt.Println("Table seller dropped successfully")
}
