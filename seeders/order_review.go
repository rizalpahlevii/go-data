package seeders

import (
	"context"
	"encoding/csv"
	"fmt"
	"main/database"
	"os"
	"time"
)

func orderReviewsSeeder() {
	fmt.Println("Importing order reviews...")
	start := time.Now()

	csvFile, err := os.Open("datasets/olist_order_reviews_dataset.csv")
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

		stmt, err := tx.Prepare("INSERT INTO order_review ( review_id,order_id, review_score, review_comment_title, review_comment_message, review_creation_date, review_answer_timestamp) VALUES (?, ?, ?, ?, ?, ?, ?)")
		if err != nil {
			panic(err)
		}
		defer stmt.Close()

		// iterate over the records start from 1 to skip the header
		for i := 1; i < len(records); i++ {
			_, err = stmt.ExecContext(ctx, records[i][0], records[i][1], records[i][2], records[i][3], records[i][4], records[i][5], records[i][6])
			if err != nil {
				panic(err)
			}
		}

		err = tx.Commit()
		if err != nil {
			panic(err)
		}

		done <- true
	}()

	<-done

	fmt.Printf("Imported %d records of order payment in %s", len(records)-1, time.Since(start))
}
