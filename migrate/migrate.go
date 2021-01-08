package migrate

import (
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/config"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/models"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/structs"
)

func MigrateKeDB() {
	config.Db.DropTableIfExists(&models.Pegawai{}, &models.Anggota{}, &structs.Buku{}, &structs.Jenis_Buku{},
		&structs.Penerbit_Buku{}, &structs.Detail_buku{}, &structs.Peminjaman{}, &structs.Detail_pinjam{}, &structs.Denda{}) //Drops the table if already exists

	config.Db.AutoMigrate(&models.Pegawai{}, &models.Anggota{}, &structs.Buku{}, &structs.Jenis_Buku{},
		&structs.Penerbit_Buku{}, &structs.Detail_buku{}, &structs.Peminjaman{}, &structs.Detail_pinjam{}, &structs.Denda{})

}
