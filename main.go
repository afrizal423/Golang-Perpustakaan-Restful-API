package main

import (
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/api"
	userController "github.com/afrizal423/Golang-Perpustakaan-Restful-API/api/v1/user"
	userService "github.com/afrizal423/Golang-Perpustakaan-Restful-API/app/business/user"
	userRepository "github.com/afrizal423/Golang-Perpustakaan-Restful-API/app/repository/user"
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

	config := configs.ServerTimeOut()
	app := fiber.New(config)

	api.RegisterPath(
		app,
		userCon,
	)

	app.Listen(":8000")

}
