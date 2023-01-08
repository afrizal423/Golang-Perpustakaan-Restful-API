package models

import "time"

type DetailPinjam struct {
	// gorm.Model
	IDDetailpinjam uint32     `gorm:"primary_key;auto_increment" json:"id"`
	IDBuku         uint32     `gorm:"column:id_buku" json:"id_buku"`
	BukuDetail     []Buku     `gorm:"foreignKey:IDBuku" json:"buku"`
	IDpeminjaman   uint32     `gorm:"column:IDPeminjaman" json:"IDPeminjaman"`
	DtPeminjaman   Peminjaman `gorm:"foreignKey:IDPeminjaman" json:"peminjaman"`
	Kondisi        string     `gorm:"size:255;not null;" json:"kondisi_buku"`
	CreatedAt      time.Time  `gorm:"not null;" json:"created_at"`
	UpdatedAt      time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Peminjaman struct {
	// gorm.Model
	IDPeminjaman   uint32    `gorm:"primary_key;auto_increment;column:IDPeminjaman" json:"id"`
	KodePeminjaman string    `gorm:"size:255;not null;unique" json:"kode_peminjaman"`
	IDAnggota      uint32    `gorm:"column:id_anggota" json:"id_anggota"`
	Anggota        []Anggota `gorm:"foreignKey:IDAnggota" json:"anggota"`

	TglPinjam     time.Time `json:"tgl_pinjam"`
	TglHrsKembali time.Time `json:"tgl_kembali"`

	Jaminan string `gorm:"size:255;not null;" json:"jaminan"`

	CreatedAt time.Time `gorm:"not null;" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Denda struct {
	// gorm.Model
	IDDenda       uint32    `gorm:"primary_key;auto_increment;column:id_data_denda" json:"id"`
	JumlahDenda   uint64    `gorm:"type:integer;not null" json:"jumlah_denda"`
	Tglpinjam     time.Time `json:"tgl_pinjam"`
	Tglhrskembali time.Time `json:"tgl_hrskembali"`
	Tglkembali    time.Time `json:"tgl_kembali"`

	KodePeminjaman   string       `gorm:"size:255;not null;unique" json:"kode_peminjaman"`
	IDKodePeminjaman []Peminjaman `gorm:"foreignKey:IDPeminjaman" json:"peminjaman"`
	IDAnggota        uint32       `gorm:"column:id_anggota" json:"id_anggota"`
	Anggota          []Anggota    `gorm:"foreignKey:IDAnggota" json:"anggota"`

	CreatedAt time.Time `gorm:"not null;" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
