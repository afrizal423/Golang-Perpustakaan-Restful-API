package BukuModels

import (
	"errors"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/structs"
	"github.com/badoux/checkmail"
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

func (u *PenulisBuku) Validasi() error {
	if u.PenulisBuku == "" {
		return errors.New("Required PenulisBuku")
	}
	if u.EmailPenulis == "" {
		return errors.New("Required EmailPenulis")
	}
	if err := checkmail.ValidateFormat(u.EmailPenulis); err != nil {
		return errors.New("Invalid Email")
	}
	return nil
}

func (u *PenulisBuku) TambahPenulisBuku(db *gorm.DB) (*PenulisBuku, error) {

	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &PenulisBuku{}, err
	}
	return u, nil
}

func (u *PenulisBuku) DetailPenulisBuku(db *gorm.DB, uid uint32) (*PenulisBuku, error) {
	penbuk := PenulisBuku{}

	err := db.Debug().Model(penbuk).Where("id_penulis = ?", uid).Find(&penbuk).Error

	if err != nil {
		return &PenulisBuku{}, err
	}
	return &penbuk, nil
}

func (u *PenulisBuku) HapusPenulisBuku(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&PenulisBuku{}).Where("id_penulis = ?", uid).Take(&PenulisBuku{}).Delete(&PenulisBuku{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (u *PenulisBuku) CariPenulisBuku(db *gorm.DB, q string) (*[]PenulisBuku, *gorm.DB) {
	// var err error
	penbuk := []PenulisBuku{}

	// tampilkan hasil update data
	qq := db.Debug().Model(&PenulisBuku{}).Where("penulis_buku LIKE ? OR alamat_penulis LIKE ? OR email_penulis LIKE ?", "%"+q+"%", "%"+q+"%", "%"+q+"%").Find(&penbuk)

	// if err != nil {
	// 	return &[]JenisBuku{}, err
	// }
	return &penbuk, qq
}

func (u *PenulisBuku) UpdatePenulisBuku(db *gorm.DB, uid uint32) (*PenulisBuku, error) {
	var err error
	penbuk := &PenulisBuku{}

	db = db.Debug().Model(penbuk).Where("id_penulis = ?", uid).Take(penbuk).UpdateColumns(
		map[string]interface{}{
			"penulis_buku":   u.PenulisBuku,
			"alamat_penulis": u.AlamatPenulis,
			"email_penulis":  u.EmailPenulis,
			"deskripsi":      u.Deskripsi,
		},
	)

	if db.Error != nil {
		return &PenulisBuku{}, db.Error
	}

	// tampilkan hasil update data
	err = db.Debug().Model(&PenulisBuku{}).Where("id_penulis = ?", uid).Take(&u).Error
	if err != nil {
		return &PenulisBuku{}, err
	}
	return penbuk, nil
}
