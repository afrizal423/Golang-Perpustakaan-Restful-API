package buku

import "github.com/afrizal423/Golang-Perpustakaan-Restful-API/app/models"

type IBukuService interface {
	CreateJenisBuku(diary models.Jenis_Buku) error
}

type IBukuRepository interface {
	CreateJenisBuku(diary models.Jenis_Buku) error
}
