package main

import (
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/configs"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/database"
)

func main() {
	db := configs.MySQLConn()
	database.Migrate(db)
	database.Seeder(db)
}
