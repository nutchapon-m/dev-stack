package routes

import (
	"backend-go/adapters/handlers"
	"backend-go/adapters/repositories"
	"backend-go/core/services"
	"backend-go/pkgs/db"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(r fiber.Router) {
	userRepositoryDB := repositories.NewUserRepositoryDB(db.MYSQL)
	userService := services.NewUserService(userRepositoryDB)
	useHandler := handlers.NewUserHandler(userService)

	r.Get("/user", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "user"})
	})
	r.Post("/user", useHandler.CreateNewUser)
}
