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

type UsersController struct {
	UsersService service.UsersService
	Enforcer     casbin.Enforcer
}

func NewUsersController(usersService *service.UsersService, enforcer *casbin.Enforcer) UsersController {
	return UsersController{UsersService: *usersService, Enforcer: *enforcer}
}

func (controller *UsersController) Route(app *fiber.App) {

	app.Post("/users/login", controller.Login)

	route := app.Group("/api", util.JWTProtected())

	route.Post("/users", util.LevelProtected("write", &controller.Enforcer), controller.Create)
	route.Get("/users", util.LevelProtected("read", &controller.Enforcer), controller.List)
}

func (controller *UsersController) Login(c *fiber.Ctx) error {
	var request model.LoginRequest
	err := c.BodyParser(&request)

	if err != nil {
		logrus.Panic(err)
	}

	token, role, err := controller.UsersService.Login(request)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(model.WebResponse{
			Code:   fiber.StatusNotFound,
			Status: fiber.ErrNotFound.Message,
			Data:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data: model.LoginResponse{
			Role:  role,
			Token: token,
		},
	})
}

func (controller *UsersController) Details(c *fiber.Ctx) error {
	return fiber.ErrBadGateway
}

func (controller *UsersController) Create(c *fiber.Ctx) error {
	var request model.CreateUsersRequest

	err := c.BodyParser(&request)
	if err != nil {
		logrus.Panic(err)
	}

	request.Id = uuid.New().String()
	util.HashPassword(&request.Password)

	response := controller.UsersService.Create(request)

	return c.Status(fiber.StatusOK).JSON(model.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *UsersController) List(c *fiber.Ctx) error {
	responses := controller.UsersService.List()
	return c.Status(fiber.StatusOK).JSON(model.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   responses,
	})
}
