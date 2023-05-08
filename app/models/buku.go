package models

import "time"

type Buku struct {
	// gorm.Model
	IDBuku          string        `gorm:"primary_key;size:26;not null;column:id_buku;" json:"id_buku"`
	ISBN            string        `gorm:"size:255;not null;unique" json:"isbn"`
	IDKategoriJenis string        `gorm:"column:id_kategori;not null;size:26;" json:"id_kategori_buku"`
	KategoriJenis   Jenis_Buku    `gorm:"foreignKey:IDKategoriJenis;" json:"kategori_buku"`
	Judul           string        `gorm:"size:255;not null;column:judul_buku" json:"judul_buku"`
	IDPenulisBuku   string        `gorm:"column:id_penulisbuku;size:26;" json:"id_penulis_buku"`
	PenulisBuku     Penulis_Buku  `gorm:"foreignKey:IDPenulisBuku" json:"penulis_buku"`
	IDPenerbitBuku  string        `gorm:"column:id_penerbitbuku;size:26;" json:"id_penerbit_buku"`
	PenerbitBuku    Penerbit_Buku `gorm:"foreignKey:IDPenerbitBuku" json:"penerbit_buku"`
	ThnTerbit       string        `gorm:"size:255;null;column:tahun_terbit" json:"tahun_terbit"`
	StokBuku        int32         `gorm:"type:integer;null;column:stok_buku" json:"stok_buku"`
	RakBuku         string        `gorm:"size:255;null;column:rak_buku" json:"rak_buku"`
	DeskripsiBuku   string        `gorm:"type:text;null;column:deskripsi_buku" json:"deskripsi_buku"`
	Gambarbuku      string        `gorm:"size:255;null;column:gambar_buku" json:"gambar_buku"`
	Kondisibuku     string        `gorm:"size:255;null;column:kondisi_buku" json:"kondisi_buku"`
	CreatedAt       time.Time     `gorm:"not null;" json:"created_at"`
	UpdatedAt       time.Time     `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Jenis_Buku struct {
	// gorm.Model
	IDJenis   string    `gorm:"primary_key;size:26;not null;" json:"id"`
	JenisBuku string    `gorm:"size:255;not null;unique" json:"jenis_buku"`
	Deskripsi string    `gorm:"type:text;null;" json:"deskripsi"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Penulis_Buku struct {
	// gorm.Model
	IDPenulis     string    `gorm:"primary_key;size:26;not null;" json:"id"`
	PenulisBuku   string    `gorm:"size:255;not null;unique" json:"penulis_buku"`
	AlamatPenulis string    `gorm:"size:255;null;" json:"alamat"`
	EmailPenulis  string    `gorm:"size:255;null;unique" json:"email_penulis"`
	Deskripsi     string    `gorm:"type:text;null;" json:"deskripsi"`
	UpdatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Penerbit_Buku struct {
	// gorm.Model
	IDPenerbit     string    `gorm:"primary_key;size:26;not null;" json:"id"`
	PenerbitBuku   string    `gorm:"size:255;not null;unique" json:"penerbit_buku"`
	AlamatPenerbit string    `gorm:"size:255;null" json:"alamat_penerbit"`
	TelpPenerbit   string    `gorm:"size:255;null" json:"telp_penerbit"`
	EmailPenerbit  string    `gorm:"size:255;null;unique" json:"email_penerbit"`
	Deskripsi      string    `gorm:"type:text;null;" json:"deskripsi_penerbit"`
	UpdatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
