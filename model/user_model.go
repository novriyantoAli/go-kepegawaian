package model

// login
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Role  string `json:"role"`
	Token string `json:"token"`
}

// create
type CreateUsersRequest struct {
	Id          string
	Username    string `json:"username"`
	Password    string `json:"password"`
	NamaLengkap string `json:"nama_lengkap"`
	Email       string `json:"email"`
	NoTelp      string `json:"no_telp"`
	Role        string `json:"role"`
}

type CreateUsersResponse struct {
	Id          string `json:"id"`
	Username    string `json:"username"`
	NamaLengkap string `json:"nama_lengkap"`
	Email       string `json:"email"`
	NoTelp      string `json:"no_telp"`
	Role        string `json:"role"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type GetUsersResponse struct {
	Id          string `json:"id"`
	Username    string `json:"username"`
	NamaLengkap string `json:"nama_lengkap"`
	Email       string `json:"email"`
	NoTelp      string `json:"no_telp"`
	Role        string `json:"role"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
