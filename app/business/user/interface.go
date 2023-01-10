package user

import "github.com/afrizal423/Golang-Perpustakaan-Restful-API/app/models"

type IUserService interface {
	Login(email string, password string) error
}

type IUserRepository interface {
	Login(email string, password string) (*models.Anggota, error)
}
