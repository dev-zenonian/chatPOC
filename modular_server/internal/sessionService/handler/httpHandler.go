package handler

import (
	"sessionService/service"

	"github.com/gofiber/fiber/v2"
)

type SessionHTTPHandler interface {
	GetSessionHandler() fiber.Handler
	GetSessionWithClientIDHandler() fiber.Handler
}

type sessionHTTPHandlerImpl struct {
	sessionService service.SessionService
}

func NewSessionHTTPHandlerImpl(sessionService service.SessionService) (SessionHTTPHandler, error) {
	hdl := &sessionHTTPHandlerImpl{
		sessionService: sessionService,
	}
	return hdl, nil
}

func (h *sessionHTTPHandlerImpl) GetSessionHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		session, err := h.sessionService.GetSessions()
		if err != nil {
			return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
				"error": true,
				"data":  err.Error(),
			})
		}
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"error": false,
			"data":  session,
		})

	}
}

func (h *sessionHTTPHandlerImpl) GetSessionWithClientIDHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		clientID := ctx.Params("clientID")
		if clientID == "" {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"data":  "clientID must not null",
			})
		}
		session, err := h.sessionService.GetsessionWithClientID(clientID)
		if err != nil {
			return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
				"error": true,
				"data":  err.Error(),
			})
		}
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"error": false,
			"data":  session,
		})

	}
}
