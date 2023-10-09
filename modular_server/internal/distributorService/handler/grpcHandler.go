package handler

import (
	"common/models"
	"common/pb"
	"context"
	"distributorService/service"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageDistributorGRPCHandler struct {
	distributor service.MessageDistributor
	pb.UnimplementedMessageDistributorServer
}

func NewMessageDistributoGRPCHandler(distributor service.MessageDistributor) (*MessageDistributorGRPCHandler, error) {
	hdl := &MessageDistributorGRPCHandler{
		distributor: distributor,
	}
	return hdl, nil
}

func (h *MessageDistributorGRPCHandler) DistributeMessage(ctx context.Context, req *pb.DistributeMessageRequest) (*pb.DistributeMessageResponse, error) {
	var (
		msgReq  = req.GetMessage()
		toID    = msgReq.GetToID()
		msgType = msgReq.GetMessageType()
	)
	fid, err := primitive.ObjectIDFromHex(msgReq.GetFromID())
	log.Printf("Received request, %v\n", req.String())
	if err != nil {
		return &pb.DistributeMessageResponse{
			Error: true,
			Data: &pb.DistributeMessageResponse_Msg{
				Msg: err.Error(),
			},
		}, nil
	}
	tid, err := primitive.ObjectIDFromHex(toID)
	if err != nil {
		return &pb.DistributeMessageResponse{
			Error: true,
			Data: &pb.DistributeMessageResponse_Msg{
				Msg: err.Error(),
			},
		}, nil
	}
	gid := primitive.NilObjectID
	mtype := models.InvidualMessage
	if msgType == models.GroupMessage.String() {
		mtype = models.GroupMessage
		gid, err = primitive.ObjectIDFromHex(msgReq.GroupID)
		if err != nil {
			return &pb.DistributeMessageResponse{
				Error: true,
				Data: &pb.DistributeMessageResponse_Msg{
					Msg: err.Error(),
				},
			}, nil
		}
	}
	mid := primitive.NewObjectID()
	msg := &models.Message{
		ID:        mid,
		FromID:    fid,
		ToID:      tid,
		GroupID:   gid,
		Data:      msgReq.GetContent(),
		Timestamp: time.Now().Unix(),
		Type:      mtype,
	}
	if err := h.distributor.SendMessage(msg); err != nil {
		return &pb.DistributeMessageResponse{
			Error: true,
			Data: &pb.DistributeMessageResponse_Msg{
				Msg: err.Error(),
			},
		}, nil
	}
	return &pb.DistributeMessageResponse{
		Error: false,
		Data: &pb.DistributeMessageResponse_Message{
			Message: msgReq,
		},
	}, nil
}
