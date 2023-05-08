package user

import (
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/api/v1/user/user_request"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/api/v1/user/user_response"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/app/business/user"
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	service user.IUserService
}

func NewUserController(service user.IUserService) *Controller {
	return &Controller{
		service,
	}
}

func (c *Controller) Login(f *fiber.Ctx) error {
	isian := new(user_request.AnggotaLoginRequest)
	if err := f.BodyParser(&isian); err != nil {
		return f.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	user, err := c.service.Login(isian.Username, isian.Password)
	if err != nil {
		return f.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	response := user_response.NewLoginResponse(user.Username, user.Token)
	return f.Status(fiber.StatusCreated).JSON(response)

}
