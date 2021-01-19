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
