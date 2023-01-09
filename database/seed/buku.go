package seed

import (
	"log"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/app/models"
	"github.com/oklog/ulid/v2"
)

var jenbuk = []models.Jenis_Buku{
	{
		IDJenis:   ulid.Make().String(),
		JenisBuku: "History",
		Deskripsi: "Ini adalah jenis buku history",
	},
	{
		IDJenis:   ulid.Make().String(),
		JenisBuku: "Cerpen",
		Deskripsi: "Ini adalah jenis buku Cerpen",
	},
	{
		IDJenis:   ulid.Make().String(),
		JenisBuku: "Kepercayaan",
		Deskripsi: "Ini adalah jenis buku Kepercayaan",
	},
	{
		IDJenis:   ulid.Make().String(),
		JenisBuku: "Politik",
		Deskripsi: "Ini adalah jenis buku Politik",
	},
	{
		IDJenis:   ulid.Make().String(),
		JenisBuku: "Ujian",
		Deskripsi: "Ini adalah jenis buku Ujian",
	},
}

var penbuk = []models.Penulis_Buku{
	{
		IDPenulis:     ulid.Make().String(),
		PenulisBuku:   "Penulis 1",
		AlamatPenulis: "Kota 1",
		EmailPenulis:  "test1@test.com",
	},
	{
		IDPenulis:     ulid.Make().String(),
		PenulisBuku:   "Penulis 2",
		AlamatPenulis: "Kota 2",
		EmailPenulis:  "test2@test.com",
	},
	{
		IDPenulis:     ulid.Make().String(),
		PenulisBuku:   "Penulis 3",
		AlamatPenulis: "Kota 3",
		EmailPenulis:  "test3@test.com",
	},
	{
		IDPenulis:     ulid.Make().String(),
		PenulisBuku:   "Penulis 4",
		AlamatPenulis: "Kota 4",
		EmailPenulis:  "test4@test.com",
	},
	{
		IDPenulis:     ulid.Make().String(),
		PenulisBuku:   "Penulis 5",
		AlamatPenulis: "Kota 5",
		EmailPenulis:  "test5@test.com",
	},
}

var PubBuk = []models.Penerbit_Buku{
	{
		IDPenerbit:     ulid.Make().String(),
		PenerbitBuku:   "Penerbit 1",
		AlamatPenerbit: "Kota 1",
		EmailPenerbit:  "publisher1@test.com",
	},
	{
		IDPenerbit:     ulid.Make().String(),
		PenerbitBuku:   "Penerbit 2",
		AlamatPenerbit: "Kota 2",
		EmailPenerbit:  "publisher2@test.com",
	},
	{
		IDPenerbit:     ulid.Make().String(),
		PenerbitBuku:   "Penerbit 3",
		AlamatPenerbit: "Kota 3",
		EmailPenerbit:  "publisher3@test.com",
	},
	{
		IDPenerbit:     ulid.Make().String(),
		PenerbitBuku:   "Penerbit 4",
		AlamatPenerbit: "Kota 4",
		EmailPenerbit:  "publisher4@test.com",
	},
	{
		IDPenerbit:     ulid.Make().String(),
		PenerbitBuku:   "Penerbit 5",
		AlamatPenerbit: "Kota 5",
		EmailPenerbit:  "publisher5@test.com",
	},
}

var dtbuku = []models.Buku{
	{
		IDBuku:          ulid.Make().String(),
		ISBN:            "INIISBNKE1",
		IDKategoriJenis: jenbuk[0].IDJenis,
		Judul:           "Ini buku ke 1",
		IDPenulisBuku:   penbuk[0].IDPenulis,
		IDPenerbitBuku:  PubBuk[0].IDPenerbit,
		ThnTerbit:       "2020",
		StokBuku:        10,
		RakBuku:         "A01",
		DeskripsiBuku:   "Ini deskripsi kali ya, bagian buku 1",
	},
	{
		IDBuku:          ulid.Make().String(),
		ISBN:            "INIISBNKE2",
		IDKategoriJenis: jenbuk[1].IDJenis,
		Judul:           "Ini buku ke 2",
		IDPenulisBuku:   penbuk[1].IDPenulis,
		IDPenerbitBuku:  PubBuk[1].IDPenerbit,
		ThnTerbit:       "2020",
		StokBuku:        10,
		RakBuku:         "A02",
		DeskripsiBuku:   "Ini deskripsi kali ya, bagian buku 2",
	},
	{
		IDBuku:          ulid.Make().String(),
		ISBN:            "INIISBNKE3",
		IDKategoriJenis: jenbuk[2].IDJenis,
		Judul:           "Ini buku ke 3",
		IDPenulisBuku:   penbuk[2].IDPenulis,
		IDPenerbitBuku:  PubBuk[2].IDPenerbit,
		ThnTerbit:       "2020",
		StokBuku:        10,
		RakBuku:         "A03",
		DeskripsiBuku:   "Ini deskripsi kali ya, bagian buku 3",
	},
	{
		IDBuku:          ulid.Make().String(),
		ISBN:            "INIISBNKE4",
		IDKategoriJenis: jenbuk[3].IDJenis,
		Judul:           "Ini buku ke 4",
		IDPenulisBuku:   penbuk[3].IDPenulis,
		IDPenerbitBuku:  PubBuk[3].IDPenerbit,
		ThnTerbit:       "2020",
		StokBuku:        10,
		RakBuku:         "A04",
		DeskripsiBuku:   "Ini deskripsi kali ya, bagian buku 4",
	},
	{
		IDBuku:          ulid.Make().String(),
		ISBN:            "INIISBNKE5",
		IDKategoriJenis: jenbuk[4].IDJenis,
		Judul:           "Ini buku ke 5",
		IDPenulisBuku:   penbuk[4].IDPenulis,
		IDPenerbitBuku:  PubBuk[4].IDPenerbit,
		ThnTerbit:       "2020",
		StokBuku:        10,
		RakBuku:         "A05",
		DeskripsiBuku:   "Ini deskripsi kali ya, bagian buku 5",
	},
}

func (s *Seeds) Seed_Buku() {
	// jenis buku
	for i, _ := range jenbuk {
		err := s.Db.Debug().Model(&models.Jenis_Buku{}).Create(&jenbuk[i]).Error
		if err != nil {
			log.Fatalf("tidak bisa seed data: %v", err)
		}
	}

	//penulis buku
	for i, _ := range penbuk {
		err := s.Db.Debug().Model(&models.Penulis_Buku{}).Create(&penbuk[i]).Error
		if err != nil {
			log.Fatalf("tidak bisa seed data: %v", err)
		}
	}

	//penerbit buku
	for i, _ := range penbuk {
		err := s.Db.Debug().Model(&models.Penerbit_Buku{}).Create(&PubBuk[i]).Error
		if err != nil {
			log.Fatalf("tidak bisa seed data: %v", err)
		}
	}

	//buku
	for i, _ := range penbuk {
		err := s.Db.Debug().Model(&models.Buku{}).Create(&dtbuku[i]).Error
		if err != nil {
			log.Fatalf("tidak bisa seed data: %v", err)
		}
	}
}
