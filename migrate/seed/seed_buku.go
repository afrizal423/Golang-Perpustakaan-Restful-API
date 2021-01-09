package seed

import (
	"log"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/config"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/structs"
)

var jenbuk = []structs.Jenis_Buku{
	structs.Jenis_Buku{
		JenisBuku: "History",
		Deskripsi: "Ini adalah jenis buku history",
	},
	structs.Jenis_Buku{
		JenisBuku: "Cerpen",
		Deskripsi: "Ini adalah jenis buku Cerpen",
	},
	structs.Jenis_Buku{
		JenisBuku: "Kepercayaan",
		Deskripsi: "Ini adalah jenis buku Kepercayaan",
	},
	structs.Jenis_Buku{
		JenisBuku: "Politik",
		Deskripsi: "Ini adalah jenis buku Politik",
	},
	structs.Jenis_Buku{
		JenisBuku: "Ujian",
		Deskripsi: "Ini adalah jenis buku Ujian",
	},
}

func Seed_JenBuk() {
	db := config.Db

	// masukkan ke db
	for i, _ := range jenbuk {
		err := db.Debug().Model(&structs.Jenis_Buku{}).Create(&jenbuk[i]).Error
		if err != nil {
			log.Fatalf("tidak bisa seed data: %v", err)
		}
	}
}
