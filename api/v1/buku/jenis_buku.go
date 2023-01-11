package buku

import (
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/api/v1/buku/buku_request"
	"github.com/gofiber/fiber/v2"
)

func (c *Controller) CreateJenisBuku(f *fiber.Ctx) error {
	jenbuk := new(buku_request.Jenis_Buku_Request)
	if err := f.BodyParser(&jenbuk); err != nil {
		return f.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return nil
}
