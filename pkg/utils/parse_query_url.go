package utils

import (
	"log"
	"net/url"

	"github.com/gofiber/fiber/v2"
)

func GetAllQueryParams(f *fiber.Ctx) url.Values {
	queryStr := f.Context().URI().QueryArgs().String()
	params, err := url.ParseQuery(queryStr)
	if err != nil {
		log.Fatal(err)
		// return
	}

	// fmt.Println("Query Params: ")
	// for key, value := range params {
	// 	fmt.Printf("  %v = %v\n", key, value)
	// }

	return params
}
