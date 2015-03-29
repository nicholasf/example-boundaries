package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/nicholasf/example/rpc/users/context/models"
)

type UsersService struct {
	users map[string]*models.User
}

func (service *UsersService) ByName(name string, user *models.User) error {
	log.Printf(fmt.Sprintf("Locating user named [%s]", name))

	user, ok := service.users[name]
	if !ok {
		return errors.New("No such user")
	}

	return nil
}

func (service *UsersService) CreateUser(name string, user *models.User) error {
	_, ok := service.users[name]

	if ok {
		return errors.New("User already exists named [" + name + "].")
	}

	service.users[name] = &models.User{Name: name}
	return nil
}

func NewUsersService() *UsersService {
	service := &UsersService{
		users: make(map[string]*models.User),
	}

	return service
}
