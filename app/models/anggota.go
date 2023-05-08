package models

import "time"

type Anggota struct {
	IDAnggota string    `gorm:"primary_key;size:26;not null;" json:"id_anggota"`
	Username  string    `gorm:"size:255;not null;unique" json:"username"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	Nama      string    `gorm:"size:255;not null;" json:"nama"`
	Sex       string    `gorm:"size:255;not null;" json:"jenis_kelamin"`
	Telp      string    `gorm:"size:255;not null;" json:"telp"`
	Alamat    string    `gorm:"type:text;not null;" json:"alamat"`
	Email     string    `gorm:"size:255;not null;unique" json:"email"`
	Deskripsi string    `gorm:"type:text;not null;" json:"deskripsi"`
	CreatedAt time.Time `gorm:"not null;" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
