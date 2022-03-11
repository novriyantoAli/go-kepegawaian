package service

import "github.com/novriyantoAli/go-kepegawaian/model"

type PegawaiService interface {
	Create(request model.CreatePegawaiRequest) (response model.CreatePegawaiResponse)

	List() (response []model.GetPegawaiResponse)
}
