package buku

import (
	"html"
	"strings"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/api/v1/buku/buku_request"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/api/v1/buku/buku_response"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func (c *Controller) GetAllPenulisBuku(f *fiber.Ctx) error {
	// for pagination
	// params := utils.GetAllQueryParams(f)
	// fmt.Println("Query Params: ")
	// for key, value := range params {
	// 	fmt.Printf("  %v = %v\n", key, value)
	// }
	cari := f.Query("q")
	if cari != "" {
		data, err := c.service.FindPenulisBuku(cari)
		if err != nil {
			return f.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}
		return f.Status(fiber.StatusOK).JSON(fiber.Map{
			"error": false,
			"msg":   "Success get data Penulis buku",
			"data":  data,
		})

	}
	data, err := c.service.GetAllPenulisBuku()
	if err != nil {
		return f.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return f.Status(fiber.StatusOK).JSON(buku_response.GetAllPenulisBukuResponse(data, "Success get data Penulis buku", false))
}

func (c *Controller) GetPenulisBukuById(f *fiber.Ctx) error {
	data, err := c.service.GetPenulisBukuById(html.EscapeString(strings.TrimSpace(f.Params("id"))))
	if err != nil {
		return f.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	response := buku_response.GetPenulisBukuByIdResponse(data, "Success get data Penulis buku", false)
	return f.Status(fiber.StatusOK).JSON(response)
}

func (c *Controller) CreatePenulisBuku(f *fiber.Ctx) error {
	jenbuk := new(buku_request.Penulis_Buku_Request)
	if err := f.BodyParser(&jenbuk); err != nil {
		return f.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	cekValid := utils.GetValidator().Struct(jenbuk)
	if cekValid != nil {
		return f.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   cekValid.Error(),
		})
	}
	data, err := c.service.CreatePenulisBuku(*jenbuk.CreatePenulisBuku())
	if err != nil {
		return f.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return f.Status(fiber.StatusCreated).JSON(buku_response.GetPenulisBukuByIdResponse(&data, "Success insert data Penulis buku", false))
}

func (c *Controller) UpdatePenulisBuku(f *fiber.Ctx) error {
	jenbuk := new(buku_request.Penulis_Buku_Request_Update)
	if err := f.BodyParser(&jenbuk); err != nil {
		return f.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	cekValid := utils.GetValidator().Struct(jenbuk)
	if cekValid != nil {
		return f.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   buku_response.ErrInvalidFormatJson,
		})
	}
	data, err := c.service.UpdatePenulisBuku(*jenbuk.UpdatePenulisBuku(jenbuk.IDPenulis))

	if err != nil {
		return f.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return f.Status(fiber.StatusCreated).JSON(buku_response.GetPenulisBukuByIdResponse(&data, "Success update data Penulis buku", false))
}

func (c *Controller) DeletePenulisBuku(f *fiber.Ctx) error {
	jenbuk := new(buku_request.Penulis_Buku_Request_Delete)
	if err := f.BodyParser(&jenbuk); err != nil {
		return f.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	cekValid := utils.GetValidator().Struct(jenbuk)
	if cekValid != nil {
		return f.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   cekValid.Error(),
		})
	}
	if err := c.service.HapusPenulisBuku(jenbuk.IDPenulis); err != nil {
		return f.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return f.Status(fiber.StatusCreated).JSON(fiber.Map{
		"msg": "Success delete Penulis buku",
	})
}
