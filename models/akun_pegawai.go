package models

import (
	"errors"
	"html"
	"log"
	"strings"
	"time"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/config"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/config/auth"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/config/helper"
)

// struct saya letakkan disini, karena tidak terlalu banyak actionnya atau tidak terllau banyak logic bisnisnya
type Pegawai struct {
	IDPegawai uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Username  string    `gorm:"size:255;not null;unique" json:"username"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (u *Pegawai) Validasi() error {
	if u.Username == "" {
		return errors.New("Username masih kosong")
	}
	if u.Password == "" {
		return errors.New("Password masih kosong")
	}
	return nil
}

func (u *Pegawai) Persiapan(action string) {
	u.IDPegawai = 0
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	if strings.ToLower(action) == "seed" || strings.ToLower(action) == "tambah" || strings.ToLower(action) == "update" {
		u.Password = helper.Hash(u.Password)
	}
}

func (u *Pegawai) ProsesLogin() (map[string]string, error) {
	var err error

	user := Pegawai{}

	err = config.Db.Debug().Model(Pegawai{}).Where("username = ?", u.Username).Take(&user).Error
	if err != nil {
		// log.Print(err)
		return nil, err
	}
	match, err := helper.VerifikasiPassword(user.Password, u.Password)
	if match == false {
		return nil, err
	}
	log.Println(match)
	return auth.BuatToken(user.IDPegawai)
}
