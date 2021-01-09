package BukuModels

import (
	"errors"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/structs"
	"github.com/jinzhu/gorm"
)

type JenisBuku struct {
	structs.Jenis_Buku
}

func (u *JenisBuku) LihatJenisBuku(db *gorm.DB) (*[]JenisBuku, error) {
	var err error
	jenbuk := []JenisBuku{}
	err = db.Debug().Model(&JenisBuku{}).Find(&jenbuk).Error
	if err != nil {
		return &[]JenisBuku{}, err
	}
	return &jenbuk, err
}

func (u *JenisBuku) Validasi() error {
	if u.JenisBuku == "" {
		return errors.New("Required JenisBuku")
	}
	if u.Deskripsi == "" {
		return errors.New("Required Deskripsi")
	}
	return nil
}

func (u *JenisBuku) TambahJenisBuku(db *gorm.DB) (*JenisBuku, error) {

	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &JenisBuku{}, err
	}
	return u, nil
}

func (u *JenisBuku) HapusJenisBuku(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&JenisBuku{}).Where("id_jenis = ?", uid).Take(&JenisBuku{}).Delete(&JenisBuku{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (u *JenisBuku) DetailJenisBuku(db *gorm.DB, uid uint32) (*JenisBuku, error) {
	jenbuk := JenisBuku{}

	err := db.Debug().Model(jenbuk).Where("id_jenis = ?", uid).Find(&jenbuk).Error

	if err != nil {
		return &JenisBuku{}, err
	}
	return &jenbuk, nil
}

func (u *JenisBuku) UpdateJenisBuku(db *gorm.DB, uid uint32) (*JenisBuku, error) {
	var err error
	jenbuk := &JenisBuku{}

	db = db.Debug().Model(jenbuk).Where("id_jenis = ?", uid).Take(jenbuk).UpdateColumns(
		map[string]interface{}{
			"jenis_buku": u.JenisBuku,
			"deskripsi":  u.Deskripsi,
		},
	)

	if db.Error != nil {
		return &JenisBuku{}, db.Error
	}
	// tampilkan hasil update data
	err = db.Debug().Model(&JenisBuku{}).Where("id_jenis = ?", uid).Take(&u).Error
	if err != nil {
		return &JenisBuku{}, err
	}
	return jenbuk, nil
}

func (u *JenisBuku) CariJenisBuku(db *gorm.DB, q string) (*[]JenisBuku, error) {
	var err error
	jenbuk := []JenisBuku{}

	// tampilkan hasil update data
	err = db.Debug().Model(&JenisBuku{}).Where("jenis_buku LIKE ? OR deskripsi LIKE ?", "%"+q+"%", "%"+q+"%").Find(&jenbuk).Error

	if err != nil {
		return &[]JenisBuku{}, err
	}
	return &jenbuk, nil
}
