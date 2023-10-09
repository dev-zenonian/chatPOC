package handler

import (
	"common/models"
	"common/pb"
	"context"
	"messageService/service"
	"time"
)

type MessageGRPCHandler struct {
	pb.MessageServiceServer
	messageService service.MessageService
}

func NewMessageGRPCHandler(messageService service.MessageService) (*MessageGRPCHandler, error) {
	hdl := &MessageGRPCHandler{
		messageService: messageService,
	}
	return hdl, nil

}

func (h *MessageGRPCHandler) SendMessage(ctx context.Context, req *pb.SendMessageRequest) (*pb.SendMessageResponse, error) {
	var (
		mtype = req.GetMessageType()
	)
	if mtype != models.InvidualMessage.String() {
		return &pb.SendMessageResponse{
			Error: true,
			Data: &pb.SendMessageResponse_Msg{
				Msg: "Must be invidual message",
			},
		}, nil

	}
	e := &models.Event{
		FromID: req.GetFromID(),
		ToID:   req.GetToID(),
		Action: models.Action{
			Type:    models.InvidualMessage,
			Content: req.GetContent(),
		},
		Timestamp: time.Now().Unix(),
	}
	msg, err := h.messageService.HandleMessage(e)
	if err != nil {
		return &pb.SendMessageResponse{
			Error: true,
			Data: &pb.SendMessageResponse_Msg{
				Msg: err.Error(),
			},
		}, nil
	}
	return &pb.SendMessageResponse{
		Error: false,
		Data: &pb.SendMessageResponse_Message{
			Message: &pb.Message{
				MessageID:   msg.ID.Hex(),
				FromID:      msg.FromID.Hex(),
				ToID:        msg.ToID.Hex(),
				MessageType: msg.Type.String(),
				Content:     msg.Data,
				GroupID:     msg.GroupID.Hex(),
				Timestamp:   msg.Timestamp,
			},
		},
	}, nil
}

func (h *MessageGRPCHandler) GetMessageWithStatus(ctx context.Context, req *pb.GetMessageWithStatusRequest) (*pb.GetMessageWithStatusResponse, error) {
	var (
		status   = req.GetStatus().Number()
		clientID = req.GetClientID()
	)

	msgs, err := h.messageService.GetMessageWithStatus(clientID, int(status))
	if err != nil {
		return &pb.GetMessageWithStatusResponse{
			Error: false,
			Msg:   err.Error(),
		}, nil
	}
	protoMessages := []*pb.Message{}
	for _, msg := range msgs {
		msgType := "invidual"
		if !msg.GroupID.IsZero() {
			msgType = "group"
		}
		protoMessages = append(protoMessages, &pb.Message{
			MessageID:   msg.ID.Hex(),
			FromID:      msg.FromID.Hex(),
			ToID:        msg.ToID.Hex(),
			MessageType: msgType,
			GroupID:     msg.GroupID.Hex(),
			Content:     msg.Data,
			Timestamp:   msg.Timestamp,
		})
	}
	return &pb.GetMessageWithStatusResponse{
		Error:    false,
		Messages: protoMessages,
	}, nil
}
