package buku_request

type Jenis_Buku_Request struct {
	// gorm.Model
	JenisBuku string `json:"jenis_buku"`
	Deskripsi string `json:"deskripsi"`
}
