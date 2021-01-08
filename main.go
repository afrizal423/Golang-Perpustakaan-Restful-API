package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/migrate"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/migrate/seed"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/routes"
)

func main() {
	// proses migrate tabel
	migrate.MigrateKeDB()
	// proses seed data admin
	seed.Load()

	r := routes.Router()
	fmt.Println("Server dijalankan pada port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
