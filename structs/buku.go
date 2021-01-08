package structs

import "time"

type Buku struct {
	// gorm.Model
	ID_buku           uint32          `gorm:"primary_key;auto_increment" json:"id"`
	ISBN              string          `gorm:"size:255;not null;unique" json:"isbn"`
	ID_Kategori_jenis uint32          `gorm:"column:id_kategori" json:"id_kategori_buku"`
	Kategori_jenis    []Jenis_Buku    `gorm:"foreignKey:ID_jenis" json:"kategori_buku"`
	Judul             string          `gorm:"size:255;not null;" json:"judul_buku"`
	ID_PenulisBuku    uint32          `gorm:"column:id_penulisbuku" json:"id_penulis_buku"`
	PenulisBuku       []Penulis_Buku  `gorm:"foreignKey:ID_penulis" json:"penulis_buku"`
	ID_penerbitBuku   uint32          `gorm:"column:id_penerbitbuku" json:"id_penerbit_buku"`
	PenerbitBuku      []Penerbit_Buku `gorm:"foreignKey:ID_penerbit" json:"penerbit_buku"`
	ThnTerbit         string          `gorm:"size:255;not null;" json:"tahun_terbit"`
	Stok_buku         uint64          `gorm:"type:integer;not null;" json:"isi_buku"`
	Rak_buku          string          `gorm:"size:255;not null;" json:"rak_buku"`
	Isi_buku          string          `gorm:"size:255;not null;" json:"isi_buku"`
	CreatedAt         time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt         time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Detail_buku struct {
	// gorm.Model
	ID_detail_buku uint32 `gorm:"primary_key;auto_increment" json:"id"`
	ID_Buku        uint32 `gorm:"column:id_buku" json:"id_buku"`
	PenerbitBuku   []Buku `gorm:"foreignKey:ID_buku" json:"buku"`
	Gambarbuku     string `gorm:"size:255;null;unique" json:"gambar_buku"`
	Kondisibuku    string `gorm:"size:255;null;unique" json:"kondisi_buku"`
	Deskripsi      string `gorm:"type:text;null;" json:"deskripsi"`
}
type Jenis_Buku struct {
	// gorm.Model
	ID_jenis  uint32 `gorm:"primary_key;auto_increment" json:"id"`
	JenisBuku string `gorm:"size:255;not null;unique" json:"isbn"`
	Deskripsi string `gorm:"type:text;null;" json:"deskripsi"`
}

type Penulis_Buku struct {
	// gorm.Model
	ID_penulis     uint32 `gorm:"primary_key;auto_increment" json:"id"`
	PenulisBuku    string `gorm:"size:255;not null;unique" json:"penulis_buku"`
	Alamat_Penulis string `gorm:"size:255;null;" json:"alamat"`
	Email_Penulis  string `gorm:"size:255;null;unique" json:"email"`
	Deskripsi      string `gorm:"type:text;not null;" json:"deskripsi"`
}

type Penerbit_Buku struct {
	// gorm.Model
	ID_penerbit     uint32 `gorm:"primary_key;auto_increment" json:"id"`
	PenerbitBuku    string `gorm:"size:255;not null;unique" json:"penerbit_buku"`
	Alamat_penerbit string `gorm:"size:255;null;unique" json:"alamat"`
	Telp_penerbit   string `gorm:"size:255;null;unique" json:"telp"`
	Email_penerbit  string `gorm:"size:255;null;unique" json:"email"`
	Deskripsi       string `gorm:"type:text;null;" json:"deskripsi"`
}
