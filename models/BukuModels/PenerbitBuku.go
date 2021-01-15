package BukuModels

import (
	"errors"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/structs"
	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
)

type PenerbitBuku struct{ structs.Penerbit_Buku }

func (u *PenerbitBuku) LihatPenerbitBuku(db *gorm.DB) (*[]PenerbitBuku, error) {
	var err error
	pubbuk := []PenerbitBuku{}
	err = db.Debug().Model(&PenerbitBuku{}).Find(&pubbuk).Error
	if err != nil {
		return &[]PenerbitBuku{}, err
	}
	return &pubbuk, err
}
func (u *PenerbitBuku) Validasi() error {
	if u.PenerbitBuku == "" {
		return errors.New("Required PenerbitBuku")
	}
	if u.EmailPenerbit == "" {
		return errors.New("Required EmailPenerbit")
	}
	if err := checkmail.ValidateFormat(u.EmailPenerbit); err != nil {
		return errors.New("Invalid Email")
	}
	return nil
}

func (u *PenerbitBuku) TambahPenerbitBuku(db *gorm.DB) (*PenerbitBuku, error) {

	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &PenerbitBuku{}, err
	}
	return u, nil
}

func (u *PenerbitBuku) DetailPenerbitBuku(db *gorm.DB, uid uint32) (*PenerbitBuku, error) {
	pubbuk := PenerbitBuku{}

	err := db.Debug().Model(pubbuk).Where("id_penerbit = ?", uid).Find(&pubbuk).Error

	if err != nil {
		return &PenerbitBuku{}, err
	}
	return &pubbuk, nil
}

func (u *PenerbitBuku) CariPenerbitBuku(db *gorm.DB, q string) (*[]PenerbitBuku, *gorm.DB) {
	// var err error
	pubbuk := []PenerbitBuku{}

	qq := db.Debug().Model(&PenerbitBuku{}).Where("penerbit_buku LIKE ? OR email_penerbit LIKE ? OR deskripsi LIKE ?", "%"+q+"%", "%"+q+"%", "%"+q+"%").Find(&pubbuk)

	return &pubbuk, qq
}
