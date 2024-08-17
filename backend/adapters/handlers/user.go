package handlers

import (
	"backend-go/core/ports"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	useSrv ports.UserService
}

func NewUserHandler(useSrv ports.UserService) userHandler {
	return userHandler{useSrv: useSrv}
}

func (h userHandler) CreateNewUser(c *fiber.Ctx) error {
	return c.SendStatus(201)
}
