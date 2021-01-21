package BukuModels

import (
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/structs"
	"github.com/jinzhu/gorm"
)

type Buku struct{ structs.Buku }

func (u *Buku) LihatSemuaBuku(db *gorm.DB) (error, *gorm.DB) {
	var err error
	book := []Buku{}
	err = db.Debug().
		Preload("KategoriJenis").Preload("PenulisBuku").
		Preload("PenerbitBuku").Find(&book).Error

	query := db.Debug().
		Preload("KategoriJenis").Preload("PenulisBuku").
		Preload("PenerbitBuku").Find(&book)

	if err != nil {
		return err, query
	}
	return err, query
}

func (u *Buku) DetailBuku(db *gorm.DB, uid uint32) (*Buku, error) {
	penbuk := Buku{}

	err := db.Debug().Model(penbuk).
		Preload("KategoriJenis").Preload("PenulisBuku").
		Preload("PenerbitBuku").
		Where("id_buku = ?", uid).Find(&penbuk).Error

	if err != nil {
		return &Buku{}, err
	}
	return &penbuk, nil
}

func (u *Buku) SearchBuku(db *gorm.DB, q string, orderby string) (*[]Buku, *gorm.DB) {
	// var err error
	jenbuk := []Buku{}
	var query *gorm.DB
	if orderby == "penulis" {
		query = db.Debug().Model(&Buku{}).
			Preload("KategoriJenis").Preload("PenulisBuku").
			Preload("PenerbitBuku").
			Joins("JOIN penulis_bukus ON bukus.id_penulisbuku = penulis_bukus.id_penulis").
			Where("penulis_bukus.penulis_buku LIKE ? OR penulis_bukus.alamat_penulis LIKE ? OR penulis_bukus.email_penulis LIKE ?", "%"+q+"%", "%"+q+"%", "%"+q+"%").
			Find(&jenbuk)
	}
	if orderby == "penerbit" {
		query = db.Debug().Model(&Buku{}).
			Preload("KategoriJenis").Preload("PenulisBuku").
			Preload("PenerbitBuku").
			Joins("JOIN penerbit_bukus ON bukus.id_penerbitbuku = penerbit_bukus.id_penerbit").
			Where("penerbit_bukus.penerbit_buku LIKE ? OR penerbit_bukus.email_penerbit LIKE ? OR penerbit_bukus.deskripsi LIKE ?", "%"+q+"%", "%"+q+"%", "%"+q+"%").
			Find(&jenbuk)

	}
	// query lama
	// Preload("KategoriJenis").Preload("PenulisBuku", "penulis_buku LIKE ? OR alamat_penulis LIKE ? OR email_penulis LIKE ?", "%"+q+"%", "%"+q+"%", "%"+q+"%").
	return &jenbuk, query
}

func (u *Buku) TambahBuku(db *gorm.DB) (*Buku, error) {

	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &Buku{}, err
	}
	return u, nil
}

func (u *Buku) HapusBuku(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&Buku{}).Where("id_buku = ?", uid).Take(&Buku{}).Delete(&Buku{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (u *Buku) UpdateBuku(db *gorm.DB, uid uint32) (*Buku, error) {
	var err error
	repBuku := &Buku{}

	db = db.Debug().Model(repBuku).Where("id_buku = ?", uid).Take(repBuku).UpdateColumns(
		map[string]interface{}{
			"isbn":            u.ISBN,
			"id_kategori":     u.IDKategoriJenis,
			"judul_buku":      u.Judul,
			"id_penulisbuku":  u.IDPenulisBuku,
			"id_penerbitbuku": u.IDPenerbitBuku,
			"tahun_terbit":    u.ThnTerbit,
			"stok_buku":       u.StokBuku,
			"rak_buku":        u.RakBuku,
			"deskripsi_buku":  u.DeskripsiBuku,
			"gambar_buku":     u.Gambarbuku,
			"kondisi_buku":    u.Kondisibuku,
		},
	)

	if db.Error != nil {
		return &Buku{}, db.Error
	}
	// tampilkan hasil update data
	err = db.Debug().Model(&Buku{}).
		Preload("KategoriJenis").Preload("PenulisBuku").
		Preload("PenerbitBuku").
		Where("id_buku = ?", uid).Take(&u).Error
	if err != nil {
		return &Buku{}, err
	}
	return repBuku, nil
}
