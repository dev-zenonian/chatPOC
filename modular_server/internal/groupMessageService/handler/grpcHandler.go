package handler

import (
	"common/models"
	"common/pb"
	"context"
	"groupMessageService/service"
	"time"
)

type GroupMessageGRPCHandler struct {
	pb.GroupMessageServiceServer
	groupMessageService service.GroupMessageService
}

func NewGroupMessageGRPCHandler(groupMessageService service.GroupMessageService) (*GroupMessageGRPCHandler, error) {
	hdl := &GroupMessageGRPCHandler{
		groupMessageService: groupMessageService,
	}
	return hdl, nil
}

func (h *GroupMessageGRPCHandler) SendGroupMessage(ctx context.Context, req *pb.SendGroupMessageRequest) (*pb.SendGroupMessageResponse, error) {
	var (
		msg = req.GetMessage()
		e   = &models.Event{
			FromID: msg.FromID,
			ToID:   msg.ToID,
			Action: models.Action{
				Type:    models.GroupMessage,
				Content: msg.Content,
			},
			Timestamp: time.Now().Unix(),
		}
	)
	if err := h.groupMessageService.HandleGroupMessage(e); err != nil {
		return &pb.SendGroupMessageResponse{
			Error: true,
			Data: &pb.SendGroupMessageResponse_Msg{
				Msg: err.Error(),
			},
		}, nil
	}
	return &pb.SendGroupMessageResponse{
		Error: false,
		Data: &pb.SendGroupMessageResponse_Message{
			Message: msg,
		},
	}, nil
}

func (h *GroupMessageGRPCHandler) GetGroupMessage(ctx context.Context, req *pb.GetGroupMessageRequest) (*pb.GetGroupMessageResponse, error) {
	var (
		groupID  = req.GetGroupID()
		clientID = req.GetClientID()
		limit    = req.GetLimit()
		offset   = req.GetOffset()
	)
	msgs, err := h.groupMessageService.GetMessageFromGroup(groupID, clientID, offset, limit)
	if err != nil {
		return &pb.GetGroupMessageResponse{
			Error: true,
			Msg:   err.Error(),
		}, nil
	}
	rspMsgs := []*pb.Message{}
	for _, msg := range msgs {
		rspMsgs = append(rspMsgs, &pb.Message{
			MessageID:   msg.ID.Hex(),
			FromID:      msg.FromID.Hex(),
			ToID:        msg.ToID.Hex(),
			GroupID:     msg.GroupID.Hex(),
			MessageType: msg.Type.String(),
			Timestamp:   msg.Timestamp,
		})
	}

	return &pb.GetGroupMessageResponse{
		Error:    false,
		Messages: rspMsgs,
	}, nil
}
