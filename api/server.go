package api

import (
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/api/v1/buku"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/api/v1/user"
	"github.com/gofiber/fiber/v2"
)

func RegisterPath(f *fiber.App,
	userCon *user.Controller,
	bukuCon *buku.Controller) {

	route := f.Group("/api")
	v1 := route.Group("/v1")

	// disabale sementara
	// v1.Use(middleware.CSRF)

	pegawaiRoutes := v1.Group("/pegawai")
	pegawaiRoutes.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "sukses",
		})
	})
	pegawaiRoutes.Post("/login", userCon.Login)
}
