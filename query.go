package main

import "main/database"

func getRevenues() {
	db := database.GetConnection()
	defer db.Close()

	//	Displays the revenue (profit) earned every day based on sales every day by looking at the price paid from datasets

}
