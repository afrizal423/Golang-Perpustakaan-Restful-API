package user

import "github.com/afrizal423/Golang-Perpustakaan-Restful-API/pkg/utils"

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

func (s *userService) Login(email string, password string) error {
	return nil
}
