package main

import (
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/configs"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/database"
)

func main() {
	configs.MySQLConn()
	database.Migrate()
	// database.Seeder(db)
}
