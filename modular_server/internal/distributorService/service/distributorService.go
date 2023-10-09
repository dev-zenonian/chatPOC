package service

import (
	"common/models"
	"common/pb"
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type MessageDistributor interface {
	SendMessage(message *models.Message) error
}

type messageDistributorImpl struct {
	sessionClient pb.SessionServiceClient
}

func NewMessageDistributorImpl(sessionClient pb.SessionServiceClient) (MessageDistributor, error) {
	service := &messageDistributorImpl{
		sessionClient: sessionClient,
	}
	return service, nil
}

func initWSClient(handler *pb.HandlerInformation) (pb.WSHandlerServiceClient, error) {
	transportOpt := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.Dial(handler.GetHandlerAddress(), transportOpt)
	if err != nil {
		return nil, err
	}
	client := pb.NewWSHandlerServiceClient(conn)
	return client, nil
}

func (s *messageDistributorImpl) SendMessage(msg *models.Message) error {
	var (
		clientID = msg.ToID.Hex()
	)
	req := &pb.GetWSHandlerWithClientIDsRequest{
		ClientID: clientID,
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	rsp, err := s.sessionClient.GetWSHandlerWithClientID(ctx, req)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("client after request from sessionservice", rsp.Handlers)
	if len(rsp.GetHandlers()) == 0 {
		// no handler onine
		// Should return error, then message service with save the message to database
		log.Println("Seen later feature")
		return fmt.Errorf("Client %v not online", msg.ToID.Hex())
	}
	for _, handler := range rsp.GetHandlers() {
		wsClient, err := initWSClient(handler)
		if err != nil {
			log.Println(err)
			return err
		}
		groupID := ""
		if gid, err := primitive.ObjectIDFromHex(msg.GroupID.Hex()); err == nil {
			groupID = gid.Hex()
		}
		msgReq := &pb.PassMessageToClientRequest{
			MessageID:      msg.ID.Hex(),
			HandlerID:      handler.HandlerID,
			HandlerAddress: handler.HandlerAddress,
			Message: &pb.Message{
				FromID:      msg.FromID.Hex(),
				ToID:        msg.ToID.Hex(),
				GroupID:     groupID,
				Content:     msg.Data,
				MessageType: msg.Type.String(),
			},
		}
		msgRsp, err := wsClient.PassMessageToClient(ctx, msgReq)
		if err != nil {
			log.Println(err)
			return err
		}
		if msgRsp.Error {
			log.Println(msgRsp.Data)
			return errors.New(msgRsp.Data)
		}
	}
	return nil
}
