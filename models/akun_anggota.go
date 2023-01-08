package models

import (
	"errors"
	"strings"
	"time"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/config/helper"
	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
)

// struct saya letakkan disini, karena tidak terlalu banyak actionnya atau tidak terllau banyak logic bisnisnya
// kemungkinan hanya CRUD data anggota saja
type Anggota struct {
	IDAnggota uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Username  string    `gorm:"size:255;not null;unique" json:"username"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	Nama      string    `gorm:"size:255;not null;" json:"nama"`
	Sex       string    `gorm:"size:255;not null;" json:"jenis_kelamin"`
	Telp      string    `gorm:"size:255;not null;" json:"telp"`
	Alamat    string    `gorm:"type:text;not null;" json:"alamat"`
	Email     string    `gorm:"size:255;not null;unique" json:"email"`
	Deskripsi string    `gorm:"type:text;not null;" json:"deskripsi"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// func persiapan untuk santiasi string berbahaya
func (u *Anggota) Persiapan(action string) {
	u.IDAnggota = 0
	u.Username = helper.EscapeStrings(u.Username)
	u.Nama = helper.EscapeStrings(u.Nama)
	u.Sex = helper.EscapeStrings(u.Sex)
	if strings.ToLower(action) == "seed" || strings.ToLower(action) == "tambah" || strings.ToLower(action) == "update" {
		u.Password = helper.Hash(u.Password)
	}
}

func (u *Anggota) ValidasiLogin(aksi string) error {
	switch strings.ToLower(aksi) {
	case "update":
		if u.Username == "" {
			return errors.New("Username masih kosong")
		}
		if u.Password == "" {
			return errors.New("Password masih kosong")
		}
		if u.Nama == "" {
			return errors.New("Nama masih kosong")
		}
		if u.Telp == "" {
			return errors.New("Telp masih kosong")
		}
		if u.Alamat == "" {
			return errors.New("Alamat masih kosong")
		}
		if u.Email == "" {
			return errors.New("Email masih kosong")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	case "login":
		if u.Username == "" {
			return errors.New("Username masih kosong")
		}
		if u.Password == "" {
			return errors.New("Password masih kosong")
		}
		return nil
	default:
		if u.Username == "" {
			return errors.New("Username masih kosong")
		}
		if u.Password == "" {
			return errors.New("Password masih kosong")
		}
		if u.Nama == "" {
			return errors.New("Nama masih kosong")
		}
		if u.Telp == "" {
			return errors.New("Telp masih kosong")
		}
		if u.Alamat == "" {
			return errors.New("Alamat masih kosong")
		}
		if u.Email == "" {
			return errors.New("Email masih kosong")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	}
}

func (u *Anggota) Register(db *gorm.DB) (*Anggota, error) {

	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &Anggota{}, err
	}
	return u, nil
}
