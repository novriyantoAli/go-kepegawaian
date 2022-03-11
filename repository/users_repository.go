package repository

import "github.com/novriyantoAli/go-kepegawaian/entity"

type UsersRepository interface {
	Insert(users entity.Users) (err error)

	FindAll() (users []entity.Users)

	Delete(id string)
}
