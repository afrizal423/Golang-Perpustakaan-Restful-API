package seed

import (
	"log"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/config"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/config/helper"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/models"
)

var admin = []models.Pegawai{
	models.Pegawai{
		Username: "admin",
		Password: helper.Hash("admin"),
	},
	models.Pegawai{
		Username: "afrizal",
		Password: helper.Hash("admin"),
	},
}

func Load() {
	db := config.Db

	// masukkan ke db
	for i, _ := range admin {
		err := db.Debug().Model(&models.Pegawai{}).Create(&admin[i]).Error
		if err != nil {
			log.Fatalf("tidak bisa seed data admin: %v", err)
		}
	}
}
