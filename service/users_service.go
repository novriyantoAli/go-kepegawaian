package service

import "github.com/novriyantoAli/go-kepegawaian/model"

type UsersService interface {
	Login(request model.LoginRequest) (token string, role string, err error)

	Create(request model.CreateUsersRequest) (response model.CreateUsersResponse)

	List() (response []model.GetUsersResponse)
}
