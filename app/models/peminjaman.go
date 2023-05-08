package models

import "time"

type DetailPinjam struct {
	// gorm.Model
	IDDetailpinjam string     `gorm:"primary_key;size:26;not null;" json:"id"`
	IDBuku         string     `gorm:"column:id_buku;size:26" json:"id_buku"`
	BukuDetail     []Buku     `gorm:"foreignKey:IDBuku" json:"buku"`
	IDpeminjaman   string     `gorm:"column:id_peminjaman;size:26" json:"id_peminjaman"`
	DtPeminjaman   Peminjaman `gorm:"foreignKey:IDPeminjaman;" json:"peminjaman"`
	Kondisi        string     `gorm:"size:255;not null;" json:"kondisi_buku"`
	CreatedAt      time.Time  `gorm:"not null;" json:"created_at"`
	UpdatedAt      time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Peminjaman struct {
	// gorm.Model
	IDPeminjaman string  `gorm:"primary_key;size:26;not null;" json:"id"`
	IDAnggota    string  `gorm:"column:id_anggota;size:26" json:"id_anggota"`
	Anggota      Anggota `gorm:"foreignKey:IDAnggota" json:"anggota"`

	TglPinjam     time.Time `json:"tgl_pinjam"`
	TglHrsKembali time.Time `json:"tgl_kembali"`

	Jaminan string `gorm:"size:255;not null;" json:"jaminan"`

	CreatedAt time.Time `gorm:"not null;" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Denda struct {
	// gorm.Model
	IDDenda          string     `gorm:"primary_key;size:26;column:id_denda;not null;" json:"id"`
	JumlahDenda      uint64     `gorm:"type:integer;not null" json:"jumlah_denda"`
	Tglpinjam        time.Time  `json:"tgl_pinjam"`
	Tglhrskembali    time.Time  `json:"tgl_hrskembali"`
	Tglkembali       time.Time  `json:"tgl_kembali"`
	IDPeminjaman     string     `gorm:"column:id_peminjaman;size:26;not null;" json:"id_peminjaman"`
	IDKodePeminjaman Peminjaman `gorm:"foreignKey:IDPeminjaman;not null;" json:"peminjaman"`
	IDAnggota        string     `gorm:"column:id_anggota;size:26;not null;" json:"id_anggota"`
	Anggota          Anggota    `gorm:"foreignKey:IDAnggota;not null;" json:"anggota"`

	CreatedAt time.Time `gorm:"not null;" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
