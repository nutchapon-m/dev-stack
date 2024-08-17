package services

import "backend-go/core/ports"

type userService struct {
	userRepo ports.UserRepository
}

func NewUserService(userRepo ports.UserRepository) ports.UserService {
	return userService{userRepo: userRepo}
}

func (s userService) NewUserAccount() error {
	panic("")
}
