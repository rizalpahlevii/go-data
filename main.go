package main

import (
	"main/migrations"
)

func main() {
	migrations.RunMigrations()
	//seeders.RunSeeders()

}
