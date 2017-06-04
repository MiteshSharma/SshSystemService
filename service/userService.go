package service

import (
	"github.com/MiteshSharma/SshSystemSetup/repository"
	"github.com/MiteshSharma/SshSystemSetup/modal"
	"github.com/satori/go.uuid"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService() *UserService {
	userService:= &UserService{}
	userService.userRepo = repository.GetUserRepository()
	return userService;
}

func (us *UserService) CreateUser(user *modal.User) (*modal.Employee, error) {
	if user.Id == "" {
		user.Id = uuid.NewV4().String()
	}
	err := us.userRepo.CreateUser(user)
	if (err != nil) {
		return nil, err
	}
	return user, nil
}

func (us *UserService) UpdateUser(user *modal.User) (*modal.Employee, error) {
	err := us.userRepo.UpdateUser(user)
	if (err != nil) {
		return nil, err
	}
	return user, nil
}

func (us *UserService) GetUsers() ([]modal.User, error) {
	users, err := us.userRepo.GetUsers()
	if (err != nil) {
		return nil, err
	}
	return users, nil
}
