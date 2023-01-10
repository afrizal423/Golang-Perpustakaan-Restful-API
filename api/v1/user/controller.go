package user

import "github.com/afrizal423/Golang-Perpustakaan-Restful-API/app/business/user"

type Controller struct {
	service user.IUserService
}

func NewUserController(service user.IUserService) *Controller {
	return &Controller{
		service,
	}
}
