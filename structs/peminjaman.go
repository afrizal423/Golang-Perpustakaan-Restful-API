package structs

import (
	"time"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/models"
)

type Peminjaman struct {
	// gorm.Model
	ID_peminjaman  uint32           `gorm:"primary_key;auto_increment;column:id_peminjaman" json:"id"`
	KodePeminjaman string           `gorm:"size:255;not null;unique" json:"kode_peminjaman"`
	ID_Anggota     uint32           `gorm:"column:id_anggota" json:"id_anggota"`
	Anggota        []models.Anggota `gorm:"foreignKey:ID_Anggota" json:"anggota"`

	Tgl_pinjam     time.Time `json:"tgl_pinjam"`
	Tgl_hrskembali time.Time `json:"tgl_kembali"`

	Jaminan string `gorm:"size:255;not null;" json:"jaminan"`

	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Detail_pinjam struct {
	// gorm.Model
	ID_detailpinjam uint32    `gorm:"primary_key;auto_increment" json:"id"`
	ID_Buku         uint32    `gorm:"column:id_buku" json:"id_buku"`
	BukuDetail      []Buku    `gorm:"foreignKey:ID_buku" json:"buku"`
	ID_peminjaman   uint32    `gorm:"column:id_peminjaman" json:"id_peminjaman"`
	DtPeminjaman    []Buku    `gorm:"foreignKey:ID_peminjaman" json:"peminjaman"`
	Kondisi         string    `gorm:"size:255;not null;" json:"kondisi_buku"`
	CreatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Denda struct {
	// gorm.Model
	ID_denda       uint32    `gorm:"primary_key;auto_increment;column:id_data_denda" json:"id"`
	JumlahDenda    uint64    `gorm:"type:integer;not null" json:"kode_peminjaman"`
	Tgl_pinjam     time.Time `json:"tgl_pinjam"`
	Tgl_hrskembali time.Time `json:"tgl_hrskembali"`
	Tgl_kembali    time.Time `json:"tgl_kembali"`

	KodePeminjaman   string           `gorm:"size:255;not null;unique" json:"kode_peminjaman"`
	IDKodePeminjaman []Peminjaman     `gorm:"foreignKey:ID_peminjaman" json:"peminjaman"`
	ID_Anggota       uint32           `gorm:"column:id_anggota" json:"id_anggota"`
	Anggota          []models.Anggota `gorm:"foreignKey:ID_Anggota" json:"anggota"`

	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
