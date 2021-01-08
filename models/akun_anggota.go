package models

import (
	"time"
)

// struct saya letakkan disini, karena tidak terlalu banyak actionnya atau tidak terllau banyak logic bisnisnya
// kemungkinan hanya CRUD data anggota saja
type Anggota struct {
	ID_Anggota uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Username   string    `gorm:"size:255;not null;unique" json:"username"`
	Password   string    `gorm:"size:100;not null;" json:"password"`
	Nama       string    `gorm:"size:255;not null;" json:"nama"`
	Sex        string    `gorm:"size:255;not null;" json:"jenis_kelamin"`
	Telp       string    `gorm:"size:255;not null;" json:"telp"`
	Alamat     string    `gorm:"type:text;not null;" json:"alamat"`
	Email      string    `gorm:"size:255;not null;" json:"email"`
	Deskripsi  string    `gorm:"type:text;not null;" json:"deskripsi"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
