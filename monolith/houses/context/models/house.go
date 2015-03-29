package models

import (
	userModels "github.com/nicholasf/example/monolith/users/context/models"
)

type House struct {
	Address string `json:"address"`
	Owner *userModels.User `json:"owner"`
}
