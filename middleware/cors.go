package middleware

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	cors "github.com/gofiber/fiber/v2/middleware/cors"
)

var CORS func(*fiber.Ctx) error = cors.New(cors.Config{
	AllowMethods: http.MethodGet,
})
