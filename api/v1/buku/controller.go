package buku

import (
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/app/business/buku"
)

type Controller struct {
	service buku.IBukuService
}

func NewBukuController(service buku.IBukuService) *Controller {
	return &Controller{
		service,
	}
}
