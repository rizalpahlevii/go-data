package main

import (
	"main/migrations"
	"main/seeders"
)

func main() {
	migrations.RunMigrations()
	seeders.RunSeeders()
}
