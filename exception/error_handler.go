package exception

import (
	"github.com/gofiber/fiber/v2"
	"github.com/novriyantoAli/go-kepegawaian/model"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	_, ok := err.(ValidationError)
	if ok {
		return ctx.JSON(model.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: fiber.ErrBadRequest.Message,
			Data:   err.Error(),
		})
	}

	return ctx.JSON(model.WebResponse{
		Code:   fiber.StatusInternalServerError,
		Status: fiber.ErrInternalServerError.Message,
		Data:   err.Error(),
	})

}
