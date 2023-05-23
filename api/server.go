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

	// pegawai route
	pegawaiRoutes := v1.Group("/pegawai")
	pegawaiRoutes.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "sukses",
		})
	})
	pegawaiRoutes.Post("/login", userCon.Login)

	//end pegawai area

	//buku area
	bukuRoute := v1.Group("/buku")

	//jenis buku route
	jenBuk := bukuRoute.Group("/jenbuk")
	jenBuk.Get("/", bukuCon.GetAllJenisBuku)
	jenBuk.Get("/:id", bukuCon.GetJenisBukuById)
	jenBuk.Post("create", bukuCon.CreateJenisBuku)
	jenBuk.Put("update", bukuCon.UpdateJenisBuku)
	jenBuk.Delete("delete", bukuCon.DeleteJenisBuku)
	// end jenis buku

	// penerbit buku area
	penbuk := bukuRoute.Group("/penbuk")
	penbuk.Get("/", bukuCon.GetAllPenerbitBuku)
	penbuk.Get("/:id", bukuCon.GetPenerbitBukuById)
	penbuk.Post("create", bukuCon.CreatePenerbitBuku)
	penbuk.Put("update", bukuCon.UpdatePenerbitBuku)
	penbuk.Delete("delete", bukuCon.DeletePenerbitBuku)
	// end penerbit buku

	// penulis buku area
	author := bukuRoute.Group("/author")
	author.Get("/", bukuCon.GetAllPenulisBuku)
	author.Get("/:id", bukuCon.GetPenulisBukuById)
	author.Post("create", bukuCon.CreatePenulisBuku)
	author.Put("update", bukuCon.UpdatePenulisBuku)
	author.Delete("delete", bukuCon.DeletePenulisBuku)
	// end penulis buku area
	//end buku area
}
