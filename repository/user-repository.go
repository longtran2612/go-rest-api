package repository

import (
	"go-rest-api/model/domain"
)

type UserRepository interface {
	Create(user domain.User) domain.User
}
