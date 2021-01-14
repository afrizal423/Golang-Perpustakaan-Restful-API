package structs

import "time"

type Buku struct {
	// gorm.Model
	IDBuku          uint32          `gorm:"primary_key;auto_increment" json:"id"`
	ISBN            string          `gorm:"size:255;not null;unique" json:"isbn"`
	IDKategoriJenis uint32          `gorm:"column:id_kategori" json:"id_kategori_buku"`
	KategoriJenis   []Jenis_Buku    `gorm:"foreignKey:IDJenis" json:"kategori_buku"`
	Judul           string          `gorm:"size:255;not null;" json:"judul_buku"`
	IDPenulisBuku   uint32          `gorm:"column:id_penulisbuku" json:"id_penulis_buku"`
	PenulisBuku     []Penulis_Buku  `gorm:"foreignKey:IDPenulis" json:"penulis_buku"`
	IDPenerbitBuku  uint32          `gorm:"column:id_penerbitbuku" json:"id_penerbit_buku"`
	PenerbitBuku    []Penerbit_Buku `gorm:"foreignKey:IDPenerbit" json:"penerbit_buku"`
	ThnTerbit       string          `gorm:"size:255;not null;" json:"tahun_terbit"`
	StokBuku        uint64          `gorm:"type:integer;not null;" json:"stok_buku"`
	RakBuku         string          `gorm:"size:255;not null;" json:"rak_buku"`
	IsiBuku         string          `gorm:"size:255;not null;" json:"isi_buku"`
	CreatedAt       time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Detail_buku struct {
	// gorm.Model
	IDDetailBuku uint32 `gorm:"primary_key;auto_increment" json:"id"`
	IDBuku       uint32 `gorm:"column:id_buku" json:"id_buku"`
	Buku         []Buku `gorm:"foreignKey:IDBuku" json:"buku"`
	Gambarbuku   string `gorm:"size:255;null;unique" json:"gambar_buku"`
	Kondisibuku  string `gorm:"size:255;null;unique" json:"kondisi_buku"`
	Deskripsi    string `gorm:"type:text;null;" json:"deskripsi"`
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
	AlamatPenerbit string `gorm:"size:255;null;unique" json:"alamat"`
	TelpPenerbit   string `gorm:"size:255;null;unique" json:"telp"`
	EmailPenerbit  string `gorm:"size:255;null;unique" json:"email"`
	Deskripsi      string `gorm:"type:text;null;" json:"deskripsi"`
}
