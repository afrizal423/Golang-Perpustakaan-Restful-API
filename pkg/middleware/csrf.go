package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/utils"
)

var CSRF = csrf.New(csrf.Config{
	KeyLookup:      "header:X-Csrf-Token",
	CookieName:     "csrf_perpus",
	CookieSameSite: "Strict",
	Expiration:     1 * time.Hour,
	KeyGenerator:   utils.UUID,
})
