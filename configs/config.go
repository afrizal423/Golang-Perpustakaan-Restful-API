package configs

import (
	"fmt"
	"os"
	"time"

	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

// memanggil konfig pada env file
func Config(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error loading file env")
	}

	return os.Getenv(key)
}

func ServerTimeOut() fiber.Config {
	readTimeoutSecondCount, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))

	return fiber.Config{
		ReadTimeout: time.Second * time.Duration(readTimeoutSecondCount),
	}
}
