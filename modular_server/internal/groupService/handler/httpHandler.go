package handler

import (
	"common/models"
	"common/utils"
	"encoding/json"
	"fmt"
	"groupService/service"

	"github.com/gofiber/fiber/v2"
)

type GroupHTTPHandler interface {
	CreateGroupHandler() fiber.Handler
	DeleteGroupHandler() fiber.Handler
	GetGroupsHandler() fiber.Handler
	GetGroupWithIDHandler() fiber.Handler
	GetClientsOfGroupHandler() fiber.Handler

	RegisterClientToGroup() fiber.Handler
	UnRegisterClientFromGroup() fiber.Handler

	AdClientsToGroupHandler() fiber.Handler
	RemovelientsToGroupHandler() fiber.Handler
}
type groupHandlerImpl struct {
	groupService service.GroupService
}

func NewGroupHTTPHandlerImpl(groupService service.GroupService) (GroupHTTPHandler, error) {
	handler := &groupHandlerImpl{
		groupService: groupService,
	}
	return handler, nil
}
func (h *groupHandlerImpl) GetGroupsHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		groups, err := h.groupService.GetGroups()
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"data":  err.Error(),
			})
		}
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"error": false,
			"data":  groups,
		})
	}
}

func (h *groupHandlerImpl) GetGroupWithIDHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		groupID := ctx.Params("id")
		if groupID == "" {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"data":  "GroupID must not be null",
			})
		}
		group, err := h.groupService.GetGroupWithID(groupID)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"data":  err.Error(),
			})
		}
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"error": false,
			"data":  group,
		})
	}
}

func (h *groupHandlerImpl) GetClientsOfGroupHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		groupID := ctx.Query("group")
		if groupID == "" {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"data":  "group id must not be null",
			})
		}
		clients, err := h.groupService.GetClientsOfGroup(groupID)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"data":  err.Error(),
			})
		}
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"error": false,
			"data":  clients,
		})
	}
}

func (h *groupHandlerImpl) RegisterClientToGroup() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		client, err := utils.GetFromFiberLocal[models.ClientModel](ctx, "user")
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"data":  err.Error(),
			})
		}
		groupID := ctx.Query("group")
		if groupID == "" {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"data":  "groupID must not null",
			})

		}
		if err := h.groupService.AddClientToGroup(groupID, client.ClientID.Hex()); err != nil {
			return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
				"error": true,
				"data":  err.Error(),
			})
		}

		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"error": false,
			"data":  fmt.Sprintf("client %v register success\n", client.ClientID.Hex()),
		})
	}
}

func (h *groupHandlerImpl) UnRegisterClientFromGroup() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		client, err := utils.GetFromFiberLocal[models.ClientModel](ctx, "user")
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"data":  err.Error(),
			})
		}
		groupID := ctx.Query("group")
		if groupID == "" {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"data":  "group ID must not null",
			})
		}
		if err := h.groupService.RemoveClientFromGroup(groupID, client.ClientID.Hex()); err != nil {
			return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
				"error": true,
				"data":  err.Error(),
			})
		}

		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"error": false,
			"data":  fmt.Sprintf("client %v unregister success\n", client.ClientID.Hex()),
		})
	}
}

type createGroupRequest struct {
	GroupName      string   `json:"group_name,omiempty"`
	GroupMembersID []string `json:"group_members_id,omiempty"`
	IsPrivate      bool     `json:"is_private,omiempty"`
}

func (h *groupHandlerImpl) CreateGroupHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		req := &createGroupRequest{}
		if err := json.Unmarshal(ctx.Body(), req); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"data":  err.Error(),
			})
		}
		client, err := utils.GetFromFiberLocal[models.ClientModel](ctx, "user")
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"data":  err.Error(),
			})
		}
		g, err := h.groupService.CreateGroup(req.GroupName, client.ClientID.Hex(), req.GroupMembersID, req.IsPrivate)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"data":  err.Error(),
			})
		}
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"error": false,
			"data":  g,
		})
	}
}

func (h *groupHandlerImpl) DeleteGroupHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		client, err := utils.GetFromFiberLocal[models.ClientModel](ctx, "user")
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"data":  err.Error(),
			})
		}
		gid := ctx.Query("group")
		if gid == "" {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"data":  "GroupID must not be nil",
			})
		}
		err = h.groupService.DeleteGroup(gid, client.ClientID.Hex())
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"data":  err.Error(),
			})
		}
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"error": false,
			"data":  "deleted",
		})
	}
}

func (h *groupHandlerImpl) AdClientsToGroupHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		req := &struct {
			ClientIDs []string `json:"clients_id,omiempty"`
			GroupID   string   `json:"group_id,omiempty"`
		}{}
		if err := json.Unmarshal(ctx.Body(), req); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"data":  err.Error(),
			})
		}
		client, err := utils.GetFromFiberLocal[models.ClientModel](ctx, "user")
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"data":  err.Error(),
			})
		}
		unvalidatedClient, err := h.groupService.AddClientsToGroupWithAdminID(req.GroupID, client.ClientID.Hex(), req.ClientIDs)
		if err != nil {
			return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
				"error": true,
				"data":  fmt.Sprintf("Unvalidated: %v, err: %v\n", unvalidatedClient, err.Error()),
			})
		}

		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"error": false,
			"data":  fmt.Sprintf("unvalidated client %v\n, Added client to group", unvalidatedClient),
		})

	}
}
func (h *groupHandlerImpl) RemovelientsToGroupHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		req := &struct {
			ClientIDs []string `json:"clients_id,omiempty"`
			GroupID   string   `json:"group_id,omiempty"`
		}{}
		if err := json.Unmarshal(ctx.Body(), req); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"data":  err.Error(),
			})
		}
		client, err := utils.GetFromFiberLocal[models.ClientModel](ctx, "user")
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"data":  err.Error(),
			})
		}
		unvalidatedClient, err := h.groupService.RemoveClientsFromGroupWithAdminID(req.GroupID, client.ClientID.Hex(), req.ClientIDs)
		if err != nil {
			return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
				"error": true,
				"data":  fmt.Sprintf("Unvalidated: %v, err: %v\n", unvalidatedClient, err.Error()),
			})
		}

		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"error": false,
			"data":  fmt.Sprintf("unvalidated client %v, Remove client from group\n", unvalidatedClient),
		})

	}

}
