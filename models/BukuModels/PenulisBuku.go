package BukuModels

import (
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/structs"
	"github.com/jinzhu/gorm"
)

type PenulisBuku struct{ structs.Penulis_Buku }

func (u *PenulisBuku) LihatPenulisBuku(db *gorm.DB) (*[]PenulisBuku, error) {
	var err error
	penbuk := []PenulisBuku{}
	err = db.Debug().Model(&PenulisBuku{}).Find(&penbuk).Error
	if err != nil {
		return &[]PenulisBuku{}, err
	}
	return &penbuk, err
}

func (u *PenulisBuku) DetailPenulisBuku(db *gorm.DB, uid uint32) (*PenulisBuku, error) {
	penbuk := PenulisBuku{}

	err := db.Debug().Model(penbuk).Where("id_penulis = ?", uid).Find(&penbuk).Error

	if err != nil {
		return &PenulisBuku{}, err
	}
	return &penbuk, nil
}
