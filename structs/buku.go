package structs

import "time"

type Buku struct {
	// gorm.Model
	IDBuku          uint32        `gorm:"primary_key;auto_increment" json:"id_buku"`
	ISBN            string        `gorm:"size:255;not null;unique" json:"isbn"`
	IDKategoriJenis uint32        `gorm:"column:id_kategori;not null;" json:"id_kategori_buku"`
	KategoriJenis   Jenis_Buku    `gorm:"foreignKey:IDJenis;" json:"kategori_buku"`
	Judul           string        `gorm:"size:255;not null;column:judul_buku" json:"judul_buku"`
	IDPenulisBuku   uint32        `gorm:"column:id_penulisbuku" json:"id_penulis_buku"`
	PenulisBuku     Penulis_Buku  `gorm:"foreignKey:IDPenulis" json:"penulis_buku"`
	IDPenerbitBuku  uint32        `gorm:"column:id_penerbitbuku" json:"id_penerbit_buku"`
	PenerbitBuku    Penerbit_Buku `gorm:"foreignKey:IDPenerbit" json:"penerbit_buku"`
	ThnTerbit       string        `gorm:"size:255;null;column:tahun_terbit" json:"tahun_terbit"`
	StokBuku        uint64        `gorm:"type:integer;null;column:stok_buku" json:"stok_buku"`
	RakBuku         string        `gorm:"size:255;null;column:rak_buku" json:"rak_buku"`
	DeskripsiBuku   string        `gorm:"type:text;null;column:deskripsi_buku" json:"deskripsi_buku"`
	Gambarbuku      string        `gorm:"size:255;null;column:gambar_buku" json:"gambar_buku"`
	Kondisibuku     string        `gorm:"size:255;null;column:kondisi_buku" json:"kondisi_buku"`
	CreatedAt       time.Time     `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       time.Time     `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Detail_buku struct {
	// gorm.Model
	IDDetailBuku uint32 `gorm:"primary_key;auto_increment" json:"id"`
	IDBuku       uint32 `gorm:"column:id_buku" json:"id_buku"`
	Buku         []Buku `gorm:"foreignKey:IDBuku" json:"buku"`
	Gambarbuku   string `gorm:"size:255;null;unique" json:"gambar_buku"`
	Kondisibuku  string `gorm:"size:255;null;unique" json:"kondisi_buku"`
}
type Jenis_Buku struct {
	// gorm.Model
	IDJenis   uint32 `gorm:"primary_key;auto_increment" json:"id"`
	JenisBuku string `gorm:"size:255;not null;unique" json:"jenis_buku"`
	Deskripsi string `gorm:"type:text;null;" json:"deskripsi"`
}

type Penulis_Buku struct {
	// gorm.Model
	IDPenulis     uint32 `gorm:"primary_key;auto_increment" json:"id"`
	PenulisBuku   string `gorm:"size:255;not null;unique" json:"penulis_buku"`
	AlamatPenulis string `gorm:"size:255;null;" json:"alamat"`
	EmailPenulis  string `gorm:"size:255;null;unique" json:"email_penulis"`
	Deskripsi     string `gorm:"type:text;null;" json:"deskripsi"`
}

type Penerbit_Buku struct {
	// gorm.Model
	IDPenerbit     uint32 `gorm:"primary_key;auto_increment" json:"id"`
	PenerbitBuku   string `gorm:"size:255;not null;unique" json:"penerbit_buku"`
	AlamatPenerbit string `gorm:"size:255;null" json:"alamat_penerbit"`
	TelpPenerbit   string `gorm:"size:255;null" json:"telp_penerbit"`
	EmailPenerbit  string `gorm:"size:255;null;unique" json:"email_penerbit"`
	Deskripsi      string `gorm:"type:text;null;" json:"deskripsi_penerbit"`
}
