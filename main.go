package main

import (
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/api"

	userController "github.com/afrizal423/Golang-Perpustakaan-Restful-API/api/v1/user"
	userService "github.com/afrizal423/Golang-Perpustakaan-Restful-API/app/business/user"
	userRepository "github.com/afrizal423/Golang-Perpustakaan-Restful-API/app/repository/user"

	bukuController "github.com/afrizal423/Golang-Perpustakaan-Restful-API/api/v1/buku"
	bukuService "github.com/afrizal423/Golang-Perpustakaan-Restful-API/app/business/buku"
	bukuRepository "github.com/afrizal423/Golang-Perpustakaan-Restful-API/app/repository/buku"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/configs"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db := configs.MySQLConn()
	// database.Migrate()
	// database.Seeder(db)

	userRepo := userRepository.NewUserRepository(db)
	userServices := userService.NewUserService(userRepo)
	userCon := userController.NewUserController(userServices)

	bukuRepo := bukuRepository.NewBukuRepository(db)
	bukuServices := bukuService.NewBukuService(bukuRepo)
	bukuCon := bukuController.NewBukuController(bukuServices)

	config := configs.ServerTimeOut()
	app := fiber.New(config)

	api.RegisterPath(
		app,
		userCon,
		bukuCon,
	)

	app.Listen(":8000")

}
