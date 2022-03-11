package service

import (
	"os"
	"time"

	"github.com/novriyantoAli/go-kepegawaian/entity"
	"github.com/novriyantoAli/go-kepegawaian/model"
	"github.com/novriyantoAli/go-kepegawaian/repository"
	"github.com/novriyantoAli/go-kepegawaian/validation"
	"github.com/sirupsen/logrus"
)

type pegawaiServiceImpl struct {
	PegawaiRepository repository.PegawaiRepository
}

func NewPegawaiService(pegawaiRepository *repository.PegawaiRepository) PegawaiService {
	return &pegawaiServiceImpl{PegawaiRepository: *pegawaiRepository}
}

func (service *pegawaiServiceImpl) List() (response []model.GetPegawaiResponse) {

	pegawais := service.PegawaiRepository.FindAll()
	for _, pegawai := range pegawais {
		response = append(response, model.GetPegawaiResponse{
			Id:          pegawai.Id,
			NamaLengkap: pegawai.NamaLengkap,
			Tmt:         pegawai.Tmt,
			CreatedAt:   pegawai.CreatedAt,
			UpdatedAt:   pegawai.UpdatedAt,
		})
	}

	return
}

func (service *pegawaiServiceImpl) Create(request model.CreatePegawaiRequest) (response model.CreatePegawaiResponse) {
	validation.ValidatePegawaiCreate(request)

	user := entity.Pegawai{
		Id:          request.Id,
		NamaLengkap: request.NamaLengkap,
		Tmt:         request.Tmt,
		CreatedAt:   time.Now().Format(os.Getenv("TIME_FORMAT")),
		UpdatedAt:   time.Now().Format(os.Getenv("TIME_FORMAT")),
	}

	err := service.PegawaiRepository.Insert(user)
	if err != nil {
		logrus.Panic(err)
	}

	response.Id = user.Id
	response.NamaLengkap = user.NamaLengkap
	response.Tmt = user.Tmt
	response.CreatedAt = user.CreatedAt
	response.UpdatedAt = user.UpdatedAt

	return response
}
