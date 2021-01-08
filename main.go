package main

import (
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/migrate"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/migrate/seed"
)

func main() {
	// proses migrate tabel
	migrate.MigrateKeDB()
	// proses seed data admin
	seed.Load()
}
