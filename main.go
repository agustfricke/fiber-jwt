package main

import (
	"log"

	"github.com/agustfricke/fiber-jwt/database"
	"github.com/agustfricke/fiber-jwt/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PUT, DELETE",
		AllowCredentials: true,
	}))

    router.SetupRoutes(app)

	database.ConnectDB()

	log.Fatal(app.Listen(":3000"))
}
