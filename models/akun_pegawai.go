package models

import (
	"time"
)

// struct saya letakkan disini, karena tidak terlalu banyak actionnya atau tidak terllau banyak logic bisnisnya
type Pegawai struct {
	IDPegawai uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Username  string    `gorm:"size:255;not null;unique" json:"username"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
