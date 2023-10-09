package service

import (
	"common/models"
	"common/pb"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"messageService/repository"
	"time"

	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageService interface {
	GetMessageFromGroup(groupID string, clientID string, offset int64, limit int64) ([]*models.Message, error)
	GetMessageWithStatus(clientID string, status int) ([]*models.Message, error)
	HandleMessage(e *models.Event) (*models.Message, error)
	UpdateMessageStatus(mid string, status int) error
}

type messageServiceImpl struct {
	messageRepo       repository.MessageRepository
	kafkaConn         *kafka.Conn
	client            pb.UserServiceClient
	distributorClient pb.MessageDistributorClient
}

func NewMessageServiceImpl(messageRepo repository.MessageRepository, userClient pb.UserServiceClient, distributeClient pb.MessageDistributorClient, kafkaConn ...*kafka.Conn) (MessageService, error) {
	service := &messageServiceImpl{
		messageRepo:       messageRepo,
		client:            userClient,
		distributorClient: distributeClient,
	}
	if len(kafkaConn) == 1 {
		service.kafkaConn = kafkaConn[0]
	}
	return service, nil
}

func (s *messageServiceImpl) GetMessageFromGroup(groupID string, clientID string, offset int64, limit int64) ([]*models.Message, error) {
	// get clients from groupService
	clients := []*models.ClientModel{}
	for _, client := range clients {
		if client.ClientID.Hex() == clientID {
			message, err := s.messageRepo.GetInvidualMessage(groupID, offset, limit)
			return message, err
		}
	}
	return nil, fmt.Errorf("Client with ID %v, not in group %v", clientID, groupID)
}

func (s *messageServiceImpl) HandleMessage(e *models.Event) (*models.Message, error) {
	return s.HandleMessageWithGRPC(e)
}

func (s *messageServiceImpl) HandleMessageWithKafka(e *models.Event) error {
	if e.Type == models.GroupMessage {

	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	req := &pb.GetUserWithIDRequest{
		UserId: e.ToID,
	}
	rsp, err := s.client.GetUserWithID(ctx, req)
	if err != nil {
		log.Println(err)
		return err
	}
	if rsp.Error == true {
		return errors.New(rsp.GetMsg())
	}
	fid, err := primitive.ObjectIDFromHex(e.FromID)
	if err != nil {
		log.Println(err)
		return err
	}
	tid, err := primitive.ObjectIDFromHex(e.ToID)
	if err != nil {
		log.Println(err)
		return err
	}
	msg := models.Message{
		ID:        primitive.NewObjectID(),
		FromID:    fid,
		ToID:      tid,
		GroupID:   primitive.NilObjectID,
		Data:      e.Content,
		Timestamp: e.Timestamp,
	}
	b, err := json.Marshal(msg)
	if err != nil {
		log.Println(err)
		return err
	}
	if _, err := s.kafkaConn.Write(b); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (s *messageServiceImpl) HandleMessageWithGRPC(e *models.Event) (*models.Message, error) {
	fid, err := primitive.ObjectIDFromHex(e.FromID)
	if err != nil {
		log.Println("HandleMessageWithGRPC::ObjectIDFromHex::FromID", err.Error())
		return nil, err
	}
	tid, err := primitive.ObjectIDFromHex(e.ToID)
	if err != nil {
		log.Println("HandleMessageWithGRPC::ObjectIDFromHex::ToID", err.Error())
		return nil, err
	}
	msg := models.Message{
		ID:        primitive.NewObjectID(),
		FromID:    fid,
		ToID:      tid,
		Status:    models.Deliveried,
		GroupID:   primitive.NilObjectID,
		Data:      e.Content,
		Timestamp: e.Timestamp,
	}
	defer func() {
		if err := s.messageRepo.SaveMessage(&msg); err != nil {
			log.Println("HandleMessageWithGRPC::SaveMessage::", err.Error())
		}
	}()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	req := &pb.GetUserWithIDRequest{
		UserId: e.ToID,
	}
	rsp, err := s.client.GetUserWithID(ctx, req)
	if err != nil {
		log.Println("HandleMessageWithGRPC::GetUserWithID::", err)
		return nil, err
	}
	if rsp.Error == true {
		log.Println("HandleMessageWithGRPC::GetUserWithIDRSP::", rsp.GetMsg())
		return nil, errors.New(rsp.GetMsg())
	}
	distributeRequest := &pb.DistributeMessageRequest{
		Message: &pb.Message{
			FromID:      e.FromID,
			ToID:        e.ToID,
			GroupID:     "",
			MessageType: e.Type.String(),
			Content:     msg.Data,
		},
	}
	distributeResponse, err := s.distributorClient.DistributeMessage(ctx, distributeRequest)
	if err != nil {
		log.Println("HandleMessageWithGRPC::DistributeMessage::Rsp", err.Error())
		return nil, err
	}

	if distributeResponse.Error {
		log.Println("HandleMessageWithGRPC::DistributeMessage::Rsp", distributeResponse.GetMsg())
		// TODO should save message to repository with unread tag
		return &msg, nil
	}
	msg.Status = models.Received
	return &msg, nil
}
func (s *messageServiceImpl) GetMessageWithStatus(clientID string, status int) ([]*models.Message, error) {
	log.Printf("GetmessageWithstatus:clientID:%v,Status:%v\n", clientID, status)
	msg, err := s.messageRepo.GetMessagesWithClientIDAndStatus(clientID, models.Status(status))
	if err != nil {
		return nil, err
	}
	log.Println("Unread message: ", msg)
	if status == int(models.Deliveried) {
		for _, m := range msg {
			s.messageRepo.UpdateMessageStatus(m.ID.Hex(), models.Received)
		}
	}
	return msg, nil
}

func (s *messageServiceImpl) UpdateMessageStatus(mid string, status int) error {
	return nil
}
