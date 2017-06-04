package repository

import (
	"github.com/MiteshSharma/SshSystemSetup/modal"
	"github.com/MiteshSharma/SshSystemSetup/repository/mongo"
)

type UserRepository interface {
	CreateUser(user *modal.User) error
	GetUsers() ([]modal.User, error)
	GetUser(id string) ([]modal.User, error)
	UpdateUser(user *modal.User) error
}

func GetUserRepository() *UserRepository {
	return mongo.NewUserRepositoryMongo();
}