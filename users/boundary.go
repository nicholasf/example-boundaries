package users

import (
"github.com/nicholasf/example/users/context/models"
)

type Users struct {}

func (u *Users) ByName(name string) (models.User, error) {
	return
}

func NewUsers() *Users {
	return &Users{}
}