package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/migrate"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/migrate/seed"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/routes"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	// cek apakah ingin migrate ulang
	if os.Getenv("MIGRATE") == "true" || os.Getenv("MIGRATE") == "True" || os.Getenv("MIGRATE") == "1" {
		// proses migrate tabel
		migrate.MigrateKeDB()
		// proses seed data admin
		seed.Load()
	}

	r := routes.Router()
	fmt.Println("Server dijalankan pada port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
