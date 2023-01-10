package api

import (
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/api/v1/user"
	"github.com/gofiber/fiber/v2"
)

func RegisterPath(f *fiber.App, userCon *user.Controller) {
	route := f.Group("/api")
	v1 := route.Group("/v1")

	anggotaRoutes := v1.Group("/anggota")
	anggotaRoutes.Get("/login", func(c *fiber.Ctx) error {
		return c.SendString("login")
	})
}
