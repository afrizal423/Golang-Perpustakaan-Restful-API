package BukuModels

import (
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/structs"
	"github.com/jinzhu/gorm"
)

type PenulisBuku struct{ structs.Penulis_Buku }

func (u *PenulisBuku) LihatPenulisBuku(db *gorm.DB) (*[]PenulisBuku, error) {
	var err error
	jenbuk := []PenulisBuku{}
	err = db.Debug().Model(&PenulisBuku{}).Find(&jenbuk).Error
	if err != nil {
		return &[]PenulisBuku{}, err
	}
	return &jenbuk, err
}
