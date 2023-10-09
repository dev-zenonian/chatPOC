package handler

import (
	"encoding/json"
	"log"
	"userService/service"

	"github.com/gofiber/fiber/v2"
)

type UserHTTPHandler interface {
	GetUserWithIDHandler() fiber.Handler
	GetUsersHandler() fiber.Handler
	CreateUserHandler() fiber.Handler
}

type userHTTPHandlerImpl struct {
	userService service.UserService
}

func NewUserHTTPHandler(userService service.UserService) (UserHTTPHandler, error) {
	hdl := &userHTTPHandlerImpl{
		userService: userService,
	}
	return hdl, nil

}

func (h *userHTTPHandlerImpl) GetUserWithIDHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userID := ctx.Query("user")
		if userID == "" {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"data":  "UserID must not be null",
			})
		}
		user, err := h.userService.GetUserWithID(userID)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"data":  err.Error(),
			})
		}
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"error": false,
			"data":  user,
		})
	}
}
func (h *userHTTPHandlerImpl) GetUsersHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		log.Println("get users")
		users, err := h.userService.GetUsers()
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"data":  err.Error(),
			})

		}
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"error": false,
			"data":  users,
		})
	}
}

type createUserRequest struct {
	UserName string `json:"user_name,omiempty"`
}

func (h *userHTTPHandlerImpl) CreateUserHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		req := &createUserRequest{}
		if err := json.Unmarshal(ctx.Body(), req); err != nil {
			return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
				"error": true,
				"data":  err.Error(),
			})
		}
		usr, err := h.userService.CreateUser(req.UserName)
		if err != nil {
			log.Println(err)
			return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
				"error": true,
				"data":  err.Error(),
			})
		}

		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"error": false,
			"data":  usr,
		})
	}
}
