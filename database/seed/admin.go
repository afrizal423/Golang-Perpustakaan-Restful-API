package seed

import (
	"log"
	"time"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/app/models"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/pkg/utils"
	"github.com/oklog/ulid/v2"
)

var hash utils.Hash
var admin = []models.Pegawai{
	{
		IDPegawai: ulid.Make().String(),
		Username:  "admin",
		Password:  hash.HashPassword("admin"),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	},
	{
		IDPegawai: ulid.Make().String(),
		Username:  "afrizal",
		Password:  hash.HashPassword("admin"),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	},
}

func (s *Seeds) Seed_admin() {
	// masukkan ke db
	for i, _ := range admin {
		err := s.Db.Model(&models.Pegawai{}).Create(&admin[i]).Error
		if err != nil {
			log.Fatalf("tidak bisa seed data admin: %v", err)
		}
	}
}
