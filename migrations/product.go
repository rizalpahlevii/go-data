package migrations

import (
	"context"
	"fmt"
	"main/database"
)

func runProduct() {
	productDown()
	productUp()
}

func productDown() {
	db := database.GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "DROP TABLE IF EXISTS product"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}
	fmt.Println("Table product dropped successfully")
}

//"product_id","product_category_name","product_name_lenght","product_description_lenght","product_photos_qty","product_weight_g","product_length_cm","product_height_cm","product_width_cm"
func productUp() {
	db := database.GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "CREATE TABLE product (product_id varchar(255) NOT NULL, product_category_name varchar(255) NULL, product_name_length int NULL, product_description_length int NULL, product_photos_qty int NULL, product_weight_g int NULL, product_length_cm int NULL, product_height_cm int NULL, product_width_cm int NULL, PRIMARY KEY (product_id))"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("Table product created successfully")
}
