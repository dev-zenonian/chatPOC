package handler

import (
	"common/models"
	"common/utils"
	"encoding/json"
	"log"
	"messageService/service"
	"time"

	"github.com/gofiber/fiber/v2"
)

type MessageHTTPHandler interface {
	PostMessageHandler() fiber.Handler
	GetMessageHandler() fiber.Handler
	UpdateMessageInChatHandler() fiber.Handler
}

type messageHTTPHandlerImpl struct {
	messageService service.MessageService
}

func NewMessageHTTPHandlerImpl(messageService service.MessageService) (MessageHTTPHandler, error) {
	hdl := &messageHTTPHandlerImpl{
		messageService: messageService,
	}
	return hdl, nil
}

type message struct {
	ToID string `json:"to_id"`
	Type string `json:"message_type"`
	Data string `json:"data"`
}

func (h *messageHTTPHandlerImpl) PostMessageHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		req := &message{}
		if err := json.Unmarshal(ctx.Body(), req); err != nil {
			log.Println("PostMessageHandler::", err)
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"data":  err.Error(),
			})
		}
		user, err := utils.GetFromFiberLocal[models.ClientModel](ctx, "user")
		if err != nil {
			log.Println("PostMessageHandler::", err)
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"data":  err.Error(),
			})
		}
		if req.Type != models.InvidualMessage.String() {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"data":  "Must redirect to group message service", //messageID
			})
		}
		msg, err := h.messageService.HandleMessage(&models.Event{
			FromID: user.ClientID.Hex(),
			ToID:   req.ToID,
			Action: models.Action{
				Type:    models.InvidualMessage,
				Content: req.Data,
			},
			Timestamp: time.Now().Unix(),
		})
		if err != nil {
			log.Println("PostMessageHandler::", err)
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"data":  err.Error(),
			})
		}
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"error": false,
			"data":  msg, //messageID
		})
	}
}

func (h *messageHTTPHandlerImpl) GetMessageHandler() fiber.Handler {
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
		messages, err := h.messageService.GetMessageFromGroup(groupID, user.ClientID.Hex(), int64(offset), int64(limit))
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

func (h *messageHTTPHandlerImpl) UpdateMessageInChatHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		chatID := ctx.Query("chatid")
		if chatID == "" {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"data":  "chatid must be not null",
			})
		}
		status := ctx.QueryInt("status", 2)
		if err := h.messageService.UpdateMessageStatus(chatID, status); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"data":  err.Error(),
			})
		}
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"error": false,
			"data":  nil,
		})

	}
}
