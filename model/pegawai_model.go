package model

// create
type CreatePegawaiRequest struct {
	Id          string
	NamaLengkap string `json:"nama_lengkap"`
	Tmt         string `json:"tmt"`
}

type CreatePegawaiResponse struct {
	Id          string `json:"id"`
	NamaLengkap string `json:"nama_lengkap"`
	Tmt         string `json:"tmt"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type GetPegawaiResponse struct {
	Id          string `json:"id"`
	NamaLengkap string `json:"nama_lengkap"`
	Tmt         string `json:"tmt"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
