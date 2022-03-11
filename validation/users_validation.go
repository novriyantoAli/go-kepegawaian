package validation

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/novriyantoAli/go-kepegawaian/exception"
	"github.com/novriyantoAli/go-kepegawaian/model"
)

func ValidateLogin(request model.LoginRequest) {
	/**
	* Username    string `json:"username"`
	* Password    string `json:"password"`
	 */

	err := validation.ValidateStruct(&request,
		validation.Field(&request.Username, validation.Required),
		validation.Field(&request.Password, validation.Required),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}

func Validate(request model.CreateUsersRequest) {

	/**
	*
	* Username    string `json:"username"`
	* Password    string `json:"password"`
	* NamaLengkap string `json:"nama_lengkap"`
	* Email       string `json:"email"`
	* NoTelp      string `json:"no_telp"`
	* Role       string `json:"role"`
	 */

	err := validation.ValidateStruct(&request,
		validation.Field(&request.Username, validation.Required),
		validation.Field(&request.Password, validation.Required),
		validation.Field(&request.NamaLengkap, validation.Required),
		validation.Field(&request.Email, validation.Required),
		validation.Field(&request.NoTelp, validation.Required),
		validation.Field(&request.Role, validation.Required),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}
