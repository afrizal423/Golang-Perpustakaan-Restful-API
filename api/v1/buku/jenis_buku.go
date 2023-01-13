package buku

import (
	"html"
	"strings"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/api/v1/buku/buku_request"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/api/v1/buku/buku_response"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func (c *Controller) GetAllJenisBuku(f *fiber.Ctx) error {
	cari := f.Query("q")
	if cari != "" {
		data, err := c.service.FindJenisBuku(cari)
		if err != nil {
			return f.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}
		return f.Status(fiber.StatusOK).JSON(data)

	}
	data, err := c.service.GetAllJenisBuku()
	if err != nil {
		return f.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return f.Status(fiber.StatusOK).JSON(data)
}

func (c *Controller) GetJenisBukuById(f *fiber.Ctx) error {
	data, err := c.service.GetJenisBukuById(html.EscapeString(strings.TrimSpace(f.Params("id"))))
	if err != nil {
		return f.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	response := buku_response.GetJenisBukuByIdResponse(data)
	return f.Status(fiber.StatusOK).JSON(response)
}

func (c *Controller) CreateJenisBuku(f *fiber.Ctx) error {
	jenbuk := new(buku_request.Jenis_Buku_Request)
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
	if err := c.service.CreateJenisBuku(*jenbuk.CreateJenisBuku()); err != nil {
		return f.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return f.Status(fiber.StatusCreated).JSON(fiber.Map{
		"msg": "Success added jenis buku",
	})
}

func (c *Controller) UpdateJenisBuku(f *fiber.Ctx) error {
	jenbuk := new(buku_request.Jenis_Buku_Request_Update)
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

	if err := c.service.UpdateJenisBuku(*jenbuk.UpdateJenisBuku(jenbuk.IDJenis)); err != nil {
		return f.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return f.Status(fiber.StatusCreated).JSON(fiber.Map{
		"msg": "Success update jenis buku",
	})
}

func (c *Controller) DeleteJenisBuku(f *fiber.Ctx) error {
	jenbuk := new(buku_request.Jenis_Buku_Request_Delete)
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
	if err := c.service.HapusJenisBuku(jenbuk.IDJenis); err != nil {
		return f.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return f.Status(fiber.StatusCreated).JSON(fiber.Map{
		"msg": "Success delete jenis buku",
	})
}
