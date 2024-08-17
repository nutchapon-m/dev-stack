package ports

import "backend-go/core/entities"

type UserRepository interface {
	CreateUser(user entities.User) error
	GetUser(username string) (*entities.User, error)
}
