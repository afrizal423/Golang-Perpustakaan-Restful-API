package database

import (
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/app/models"
	"gorm.io/gorm"
)

// proses migrasi
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Pegawai{}, &models.DetailPinjam{},
		&models.Buku{}, &models.Denda{}, &models.Peminjaman{}, &models.Anggota{})

	//&models.DetailPinjam{},
}
