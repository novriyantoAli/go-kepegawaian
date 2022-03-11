package repository

import "github.com/novriyantoAli/go-kepegawaian/entity"

type PegawaiRepository interface {
	Insert(pegawais entity.Pegawai) (err error)

	FindAll() (pegawai []entity.Pegawai)

	Delete(id string)
}
