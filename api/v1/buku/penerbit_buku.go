package buku

import (
	"html"
	"strings"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/api/v1/buku/buku_request"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/api/v1/buku/buku_response"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func (c *Controller) GetAllPenerbitBuku(f *fiber.Ctx) error {
	// for pagination
	// params := utils.GetAllQueryParams(f)
	// fmt.Println("Query Params: ")
	// for key, value := range params {
	// 	fmt.Printf("  %v = %v\n", key, value)
	// }
	cari := f.Query("q")
	if cari != "" {
		data, err := c.service.FindPenerbitBuku(cari)
		if err != nil {
			return f.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}
		return f.Status(fiber.StatusOK).JSON(fiber.Map{
			"error": false,
			"msg":   "Success get data penerbit buku",
			"data":  data,
		})

	}
	data, err := c.service.GetAllPenerbitBuku()
	if err != nil {
		return f.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return f.Status(fiber.StatusOK).JSON(buku_response.GetAllPenerbitBukuResponse(data, "Success get data penerbit buku", false))
}

func (c *Controller) GetPenerbitBukuById(f *fiber.Ctx) error {
	data, err := c.service.GetPenerbitBukuById(html.EscapeString(strings.TrimSpace(f.Params("id"))))
	if err != nil {
		return f.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	response := buku_response.GetPenerbitBukuByIdResponse(data, "Success get data penerbit buku", false)
	return f.Status(fiber.StatusOK).JSON(response)
}

func (c *Controller) CreatePenerbitBuku(f *fiber.Ctx) error {
	jenbuk := new(buku_request.Penerbit_Buku_Request)
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
	data, err := c.service.CreatePenerbitBuku(*jenbuk.CreatePenerbitBuku())
	if err != nil {
		return f.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return f.Status(fiber.StatusCreated).JSON(buku_response.GetPenerbitBukuByIdResponse(&data, "Success insert data penerbit buku", false))
}

func (c *Controller) UpdatePenerbitBuku(f *fiber.Ctx) error {
	jenbuk := new(buku_request.Penerbit_Buku_Request_Update)
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
	data, err := c.service.UpdatePenerbitBuku(*jenbuk.UpdatePenerbitBuku(jenbuk.IDPenerbit))

	if err != nil {
		return f.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return f.Status(fiber.StatusCreated).JSON(buku_response.GetPenerbitBukuByIdResponse(&data, "Success update data penerbit buku", false))
}

func (c *Controller) DeletePenerbitBuku(f *fiber.Ctx) error {
	jenbuk := new(buku_request.Penerbit_Buku_Request_Delete)
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
	if err := c.service.HapusPenerbitBuku(jenbuk.IDPenerbit); err != nil {
		return f.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return f.Status(fiber.StatusCreated).JSON(fiber.Map{
		"msg": "Success delete penerbit buku",
	})
}
