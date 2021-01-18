package BukuModels

import (
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/structs"
	"github.com/jinzhu/gorm"
)

type Buku struct{ structs.Buku }

func (u *Buku) LihatSemuaBuku(db *gorm.DB) (*[]Buku, error) {
	var err error
	book := []Buku{}
	err = db.Debug().
		Preload("KategoriJenis").Preload("PenulisBuku").
		Preload("PenerbitBuku").Find(&book).Error
	if err != nil {
		return &[]Buku{}, err
	}
	return &book, err
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
