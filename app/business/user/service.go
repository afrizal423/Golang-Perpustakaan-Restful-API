package user

import (
	"errors"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/api/v1/user/user_response"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/pkg/utils"
)

type userService struct {
	repository IUserRepository
	hashing    utils.Hash
}

func NewUserService(repository IUserRepository) IUserService {
	return &userService{
		repository,
		utils.Hash{},
	}
}

func (s *userService) Login(username string, password string) (*user_response.LoginResponse, error) {
	pass := s.hashing.HashPassword(password)
	data, err := s.repository.GetDataByUsername(username)
	if err != nil {
		return nil, errors.New("username atau password salah")
	}
	cek, err := s.hashing.VerifikasiPassword(password, data.Password)
	if err != nil {
		return nil, err
	}
	if cek {
		return &user_response.LoginResponse{
			Username: data.Username,
			Token:    pass,
		}, nil
	} else {
		return nil, errors.New("username atau password salah")
	}

}
