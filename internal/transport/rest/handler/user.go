package handler

import (
	"HomeworkWebProject/internal/model"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type UserService interface {
	GetAll() []*model.User
}
type UserHandler struct {
	service UserService
}

func (handler *UserHandler) InitRoutes(app *fiber.App) {
	app.Get("/users", handler.GetAll)
}

func NewUserHandler(service UserService) *UserHandler {
	return &UserHandler{service: service}
}

// GetAll 		godoc
// @Summary 	Get all users
// @Description Get list of users
// @Tags 		users
// @Produce 	json
// @Success 	200 {object} 		[]model.User
// @Router 		/users [get]
func (handler *UserHandler) GetAll(ctx *fiber.Ctx) error {
	user := handler.service.GetAll()

	return ctx.Status(http.StatusOK).JSON(
		fiber.Map{
			"users": user,
		})
}
