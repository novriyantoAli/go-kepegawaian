package controller

import (
	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/novriyantoAli/go-kepegawaian/model"
	"github.com/novriyantoAli/go-kepegawaian/service"
	"github.com/novriyantoAli/go-kepegawaian/util"
	"github.com/sirupsen/logrus"
)

type PegawaiController struct {
	PegawaiService service.PegawaiService
	Enforcer casbin.Enforcer
}

func NewPegawaiController(pegawaiService *service.PegawaiService, enforcer *casbin.Enforcer) PegawaiController {
	return PegawaiController{PegawaiService: *pegawaiService, Enforcer: *enforcer}
}

func (controller *PegawaiController) Route(app *fiber.App) {
	route := app.Group("/api", util.JWTProtected())

	route.Post("/pegawai", util.LevelProtected("write", &controller.Enforcer), controller.Create)
	route.Get("/pegawai", util.LevelProtected("read", &controller.Enforcer), controller.List)
}

func (controller *PegawaiController) Create(c *fiber.Ctx) error {
	var request model.CreatePegawaiRequest

	err := c.BodyParser(&request)
	if err != nil {
		logrus.Panic(err)
	}

	request.Id = uuid.New().String()
	response := controller.PegawaiService.Create(request)

	return c.Status(fiber.StatusOK).JSON(model.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *PegawaiController) List(c *fiber.Ctx) error {
	responses := controller.PegawaiService.List()
	return c.Status(fiber.StatusOK).JSON(model.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   responses,
	})
}
