package migrations

import (
	"context"
	"fmt"
	"main/database"
)

func runCustomer() {
	customerDown()
	customerUp()
}

//"customer_id","customer_unique_id","customer_zip_code_prefix","customer_city","customer_state"
//"06b8999e2fba1a1fbc88172c00ba8bc7","861eff4711a542e4b93843c6dd7febb0","14409",franca,SP
func customerDown() {
	db := database.GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := `DROP TABLE IF EXISTS customer`
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}
	fmt.Println("Table customer dropped successfully")

}

func customerUp() {
	db := database.GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := `CREATE TABLE customer (
		customer_id VARCHAR(255) NOT NULL,
		customer_unique_id VARCHAR(255) NOT NULL,
		customer_zip_code_prefix VARCHAR(255) NOT NULL,
		customer_city VARCHAR(255) NOT NULL,
		customer_state VARCHAR(255) NOT NULL,
		PRIMARY KEY (customer_id)
	)`
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}
	fmt.Println("Table customer created successfully")
}
