package database

import (
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/database/seed"
	"gorm.io/gorm"
)

// proses migrasi
func Seeder(dbx *gorm.DB) {
	s := seed.Seeds{
		Db: dbx,
	}
	s.Seed_admin()
	s.Seed_Buku()
}
