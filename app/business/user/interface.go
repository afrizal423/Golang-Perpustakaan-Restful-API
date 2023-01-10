package user

import (
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/api/v1/user/user_response"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/app/models"
)

type IUserService interface {
	Login(username string, password string) (*user_response.LoginResponse, error)
}

type IUserRepository interface {
	GetDataByUsername(username string) (*models.Pegawai, error)
}
