package repositories

import (
	"backend-go/core/entities"
	"backend-go/core/ports"

	"gorm.io/gorm"
)

type userRepositoryDB struct {
	db *gorm.DB
}

func NewUserRepositoryDB(db *gorm.DB) ports.UserRepository {
	return userRepositoryDB{db: db}
}

func (r userRepositoryDB) CreateUser(user entities.User) error {
	panic("")
}

func (r userRepositoryDB) GetUser(username string) (*entities.User, error) {
	panic("")
}
