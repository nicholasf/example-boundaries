package users

import (
	"github.com/nicholasf/example/distributed/users/context/models"
	"errors"
	"fmt"
	"log"
)

type Users struct {
	users map[string]*models.User
}

func (u *Users) ByName(name string) (user *models.User, err error) {
	log.Printf(fmt.Sprintf("Locating user named [%s]", name))
	user, ok := u.users[name]

	if !ok {
		err = errors.New("No such user")
	}

	return
}

func (u *Users) CreateUser(name string) (err error) {
	_, ok := u.users[name]

	if ok {
		err = errors.New("User already exists named [" + name + "].")
		return
	}

	u.users[name] = &models.User{ Name: name }
	return
}

func NewUsers() *Users {
	return &Users{
		users: make(map[string]*models.User),
	}
}