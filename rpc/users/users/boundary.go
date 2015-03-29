package users

import (
	"net/rpc"

	"github.com/nicholasf/example/rpc/users/context/models"
)

type Users struct {
	serviceAddress string
}

func (u *Users) ByName(name string) (*models.User, error) {
	user := &models.User{}

	err := u.callRemote("UsersService.ByName", name, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *Users) CreateUser(name string) error {
	user := &models.User{}

	err := u.callRemote("UsersService.CreateUser", name, user)
	if err != nil {
		return err
	}

	return nil
}

func (u *Users) callRemote(command string, arg1 interface{}, arg2 interface{}) error {
	client, err := rpc.DialHTTP("tcp", u.serviceAddress)
	if err != nil {
		return err
	}

	defer client.Close()

	err = client.Call(command, arg1, arg2)

	if err != nil {
		return err
	}

	return nil
}

func NewUsers(serviceAddress string) *Users {
	return &Users{
		serviceAddress: serviceAddress,
	}
}
