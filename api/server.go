package api

import (
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/api/v1/user"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterPath(f *fiber.App, userCon *user.Controller) {
	route := f.Group("/api")
	v1 := route.Group("/v1")

	v1.Use(middleware.CSRF)

	anggotaRoutes := v1.Group("/anggota")
	anggotaRoutes.Get("/login", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"name": "John",
			"age":  33,
		})
	})
	anggotaRoutes.Post("/login", func(c *fiber.Ctx) error {
		isian := new(user.AnggotaLoginRequest)
		if err := c.BodyParser(&isian); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"username": isian.Username,
			"password": isian.Password,
		})
	})
}
