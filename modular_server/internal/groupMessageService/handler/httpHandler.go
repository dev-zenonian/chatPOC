package handler

import (
	"common/models"
	"common/utils"
	"encoding/json"
	"groupMessageService/service"
	"time"

	"github.com/gofiber/fiber/v2"
)

type GroupMessageHandler interface {
	PostGroupMessageHandler() fiber.Handler
	GetGroupMessageHandler() fiber.Handler
}

type groupMessageHandlerImpl struct {
	groupMessageService service.GroupMessageService
}

func NewGroupMessageHTTPHandler(groupMessageService service.GroupMessageService) (GroupMessageHandler, error) {
	hdl := &groupMessageHandlerImpl{
		groupMessageService: groupMessageService,
	}
	return hdl, nil
}

type message struct {
	ToID string `json:"to_id"`
	Type string `json:"message_type"`
	Data string `json:"data"`
}

func (h *groupMessageHandlerImpl) PostGroupMessageHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		req := &message{}
		if err := json.Unmarshal(ctx.Body(), req); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"data":  err.Error(),
			})
		}

		user, err := utils.GetFromFiberLocal[models.ClientModel](ctx, "user")
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"data":  err.Error(),
			})
		}
		if req.Type != models.GroupMessage.String() {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"data":  "Must redirect to invidual message service", //messageID
			})
		}
		if err := h.groupMessageService.HandleGroupMessage(&models.Event{
			FromID: user.ClientID.Hex(),
			ToID:   req.ToID,
			Action: models.Action{
				Type:    models.GroupMessage,
				Content: req.Data,
			},
			Timestamp: time.Now().Unix(),
		}); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"data":  err.Error(),
			})
		}
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"error": false,
			"data":  "deliveried", //messageID
		})
	}
}

func (h *groupMessageHandlerImpl) GetGroupMessageHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var (
			groupID = ctx.Query("group")
			offset  = ctx.QueryInt("offset")
			limit   = ctx.QueryInt("limit")
		)
		if offset < 0 || limit <= 0 {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"data":  "Offset or limit must be positive number",
			})
		}
		if groupID == "" {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"data":  "GroupID must not null",
			})
		}
		user, err := utils.GetFromFiberLocal[models.ClientModel](ctx, "user")
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"data":  err.Error(),
			})
		}
		messages, err := h.groupMessageService.GetMessageFromGroup(groupID, user.ClientID.Hex(), int64(offset), int64(limit))
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"data":  err.Error(),
			})
		}
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"error": false,
			"data":  messages,
		})

	}
}
