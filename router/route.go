package router

import (
	"github.com/agustfricke/fiber-jwt/handler"
	"github.com/agustfricke/fiber-jwt/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	app.Group("/", logger.New())

	app.Post("login", handler.SignInUser)
	app.Post("register", handler.SignUpUser)
	app.Post("logout", middleware.Auth, handler.LogoutUser)
    app.Get("me/:username", middleware.Auth, handler.GetMe)

}
