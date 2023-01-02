package migrations

import (
	"context"
	"fmt"
	"main/database"
)

func runTranslation() {
	translationDown()
	translationUp()
}

func translationDown() {
	db := database.GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := `DROP TABLE IF EXISTS translation`
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}
	fmt.Println("Table translation dropped successfully")
}

//product_category_name,product_category_name_english
func translationUp() {
	db := database.GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := `CREATE TABLE translation (
		product_category_name varchar(255) NOT NULL,
		product_category_name_english varchar(255) NOT NULL
		)`
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}
	fmt.Println("Table translation created successfully")
}
