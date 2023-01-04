package migrations

import (
	"context"
	"fmt"
	"main/configurations"
	"main/database"
)

func runOrderReview() {
	orderReviewDown()
	orderReviewUp()
}

//"review_id","order_id","review_score","review_comment_title","review_comment_message","review_creation_date","review_answer_timestamp"
func orderReviewDown() {
	db := database.GetConnection()
	defer db.Close()
	ctx := context.Background()
	// Drop table if exists
	dropQuery := "DROP TABLE IF EXISTS " + configurations.Configuration().OrderReviewTableName
	_, err := db.ExecContext(ctx, dropQuery)
	if err != nil {
		panic(err)
	}
	fmt.Println("Table order_review dropped successfully")
}

func orderReviewUp() {
	db := database.GetConnection()
	defer db.Close()
	ctx := context.Background()

	// Create table order_review
	script := "CREATE TABLE " + configurations.Configuration().OrderReviewTableName + " (review_id VARCHAR(50), order_id VARCHAR(50), review_score INT, review_comment_title VARCHAR(255) NULL, review_comment_message VARCHAR(255) NULL, review_creation_date DATETIME, review_answer_timestamp TIMESTAMP)"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}
	fmt.Println("Table order_review created successfully")
}
