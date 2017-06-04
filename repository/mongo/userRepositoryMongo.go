package mongo

import (
	"github.com/MiteshSharma/SshSystemSetup/modal"
	"github.com/MiteshSharma/SshSystemSetup/repository/mongo/database"
	"gopkg.in/mgo.v2/bson"
)

type UserRepositoryMongo struct {
}

func NewUserRepositoryMongo() *UserRepositoryMongo {
	userRepo:= &UserRepositoryMongo{}
	return userRepo;
}

func (urm UserRepositoryMongo) CreateUser(user *modal.User) error {
	db := database.GetDatabaseManager()
	if err := db.Create("users", user); err != nil {
		return err
	}

	return nil
}

func (urm UserRepositoryMongo) GetUsers() ([]modal.User, error) {
	db := database.GetDatabaseManager()

	result := []modal.User{}
	query := &bson.M{}
	if err := db.GetAllByQuery("users", query, &result); err != nil {
		return result, err
	}

	return result, nil
}

func (urm UserRepositoryMongo) GetUser(id string) ([]modal.User, error) {
	db := database.GetDatabaseManager()

	user := []modal.User{}
	if err := db.Get("users", id, &user); err != nil {
		return user, err
	}

	return user, nil
}

func (urm UserRepositoryMongo) UpdateUser(user *modal.User) error {
	db := database.GetDatabaseManager()
	if err := db.Save("users", user.Id, user); err != nil {
		return err
	}

	return nil
}
