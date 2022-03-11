package service

import (
	"errors"
	"os"
	"time"

	"github.com/novriyantoAli/go-kepegawaian/entity"
	"github.com/novriyantoAli/go-kepegawaian/model"
	"github.com/novriyantoAli/go-kepegawaian/repository"
	"github.com/novriyantoAli/go-kepegawaian/util"
	"github.com/novriyantoAli/go-kepegawaian/validation"
	"github.com/sirupsen/logrus"
)

type usersServiceImpl struct {
	UsersRepository repository.UsersRepository
}

func NewUsersService(usersRepository *repository.UsersRepository) UsersService {
	return &usersServiceImpl{UsersRepository: *usersRepository}
}

func (service *usersServiceImpl) List() (response []model.GetUsersResponse) {

	users := service.UsersRepository.FindAll()
	for _, user := range users {
		response = append(response, model.GetUsersResponse{
			Id:          user.Id,
			Username:    user.Username,
			NamaLengkap: user.NamaLengkap,
			Email:       user.Email,
			NoTelp:      user.NoTelp,
			Role:        user.Role,
			CreatedAt:   user.CreatedAt,
			UpdatedAt:   user.UpdatedAt,
		})
	}

	return
}

func (service *usersServiceImpl) Login(request model.LoginRequest) (token string, role string, err error) {
	validation.ValidateLogin(request)

	user := entity.Users{}
	users := service.UsersRepository.FindAll()
	for _, usr := range users {
		if usr.Username == request.Username {
			user = usr
			break
		}
	}

	if user == (entity.Users{}) {
		logrus.Warning("users not found")
	}

	if request.Username == user.Username {
		if isOk := util.ComparePassword(user.Password, request.Password); isOk {
			token, err := util.GenerateNewAccessToken(user.Id, user.Username)
			if err != nil {
				logrus.Panic(err)
			}

			return token, user.Role, nil
		}
	}

	return "", "", errors.New("invalid username or password")
}

func (service *usersServiceImpl) Create(request model.CreateUsersRequest) (response model.CreateUsersResponse) {
	validation.Validate(request)

	user := entity.Users{
		Id:          request.Id,
		Username:    request.Username,
		Password:    request.Password,
		NamaLengkap: request.NamaLengkap,
		Email:       request.Email,
		NoTelp:      request.NoTelp,
		Role:        request.Role,
		CreatedAt:   time.Now().Format(os.Getenv("TIME_FORMAT")),
		UpdatedAt:   time.Now().Format(os.Getenv("TIME_FORMAT")),
	}

	err := service.UsersRepository.Insert(user)
	if err != nil {
		logrus.Panic(err)
	}

	response.Id = user.Id
	response.Username = user.Username
	response.NamaLengkap = user.NamaLengkap
	response.Email = user.Email
	response.NoTelp = user.NoTelp
	response.Role = user.Role
	response.CreatedAt = user.CreatedAt
	response.UpdatedAt = user.UpdatedAt

	return response
}
