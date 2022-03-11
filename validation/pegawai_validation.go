package validation

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/novriyantoAli/go-kepegawaian/exception"
	"github.com/novriyantoAli/go-kepegawaian/model"
)

func ValidatePegawaiCreate(request model.CreatePegawaiRequest) {

	/**
	* NamaLengkap string `json:"nama_lengkap"`
	* Tmt       string `json:"tmt"`
	 */

	err := validation.ValidateStruct(&request,
		validation.Field(&request.NamaLengkap, validation.Required),
		validation.Field(&request.Tmt, validation.Required),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}
