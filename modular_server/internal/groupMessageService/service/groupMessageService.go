package service

import (
	"common/models"
	"common/pb"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"groupMessageService/repository"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GroupMessageService interface {
	HandleGroupMessage(event *models.Event) error
	GetMessageFromGroup(groupID string, ClientID string, offset int64, limit int64) ([]*models.Message, error)
}
type groupMessageServiceImpl struct {
	kafkaConn        *kafka.Conn
	groupClient      pb.GroupServiceClient
	groupMessageRepo repository.GroupMessageRepository
	distributor      pb.MessageDistributorClient
}

func NewGroupMessageServiceImpl(groupMessageRepo repository.GroupMessageRepository, groupClient pb.GroupServiceClient, distributor pb.MessageDistributorClient, kafkaConn ...*kafka.Conn) (GroupMessageService, error) {
	service := &groupMessageServiceImpl{
		groupMessageRepo: groupMessageRepo,
		groupClient:      groupClient,
		distributor:      distributor,
	}
	if len(kafkaConn) == 1 {
		service.kafkaConn = kafkaConn[0]
	}
	return service, nil
}
func (s *groupMessageServiceImpl) HandleGroupMessage(event *models.Event) error {
	return s.HandleGroupMessageWithGRPC(event)

}

func (s *groupMessageServiceImpl) HandleGroupMessageWithGRPC(event *models.Event) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	req := &pb.GetGroupWithIDRequest{
		GroupId: event.ToID,
	}

	rsp, err := s.groupClient.GetGroupWithID(ctx, req)
	if err != nil {
		return err
	}
	if rsp.Error {
		return fmt.Errorf("%v", rsp.GetMsg())
	}
	var (
		group = rsp.GetGroup()
		cids  = group.GetClientIDs()
	)
	err = fmt.Errorf("Client %v not in group\n", event.FromID)
	for _, cid := range cids {
		if cid == event.FromID {
			err = nil
		}
	}
	if err != nil {
		return err
	}
	fid, err := primitive.ObjectIDFromHex(event.FromID)
	if err != nil {
		return err
	}
	groupID, err := primitive.ObjectIDFromHex(event.ToID)
	if err != nil {
		return err
	}
	msg := &models.Message{
		ID:        primitive.NewObjectID(),
		FromID:    fid,
		GroupID:   groupID,
		Status:    models.Deliveried,
		Timestamp: event.Timestamp,
		Data:      event.ToID,
	}
	defer func() {
		s.groupMessageRepo.SaveMessageInGroups(msg)
	}()
	// TODO: Seen later for group message
	for _, cid := range cids {
		if cid == event.FromID {
			continue
		}

		msg := &pb.Message{
			MessageID:   msg.ID.Hex(),
			FromID:      msg.FromID.Hex(),
			ToID:        cid,
			MessageType: event.Action.Type.String(),
			Content:     msg.Data,
			GroupID:     msg.GroupID.Hex(),
			Timestamp:   msg.Timestamp,
		}

		distributeRequest := &pb.DistributeMessageRequest{
			Message: msg,
		}

		distributeResponse, err := s.distributor.DistributeMessage(ctx, distributeRequest)
		if err != nil {
			log.Println("HandleMessageWithGRPC::DistributeMessage::Rsp", err.Error())
		}

		if distributeResponse.Error {
			log.Println("HandleMessageWithGRPC::DistributeMessage::Rsp", distributeResponse.GetMsg())
			// TODO should save message to repository with unread tag
		}
	}
	return nil
}

func (s *groupMessageServiceImpl) HandleGroupMessageWithKafka(event *models.Event) error {
	// log.Printf("Broadcast event with type: %v, data:  %v, from: %v, to: %v\n", event.Type, event.Content, event.FromID, event.ToID)
	// Send to kafka
	msgCh := make(chan kafka.Message, 10)
	doneCh := make(chan struct{}, 10)
	errCh := make(chan error)
	go func() {
		cid, err := primitive.ObjectIDFromHex(event.FromID)
		if err != nil {
			errCh <- err
			return
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		req := &pb.GetGroupWithIDRequest{
			GroupId: event.ToID,
		}
		rsp, err := s.groupClient.GetGroupWithID(ctx, req)
		if err != nil {
			errCh <- err
			return
		}
		if rsp.Error == true {
			log.Println(rsp.GetMsg())
			errCh <- errors.New(rsp.GetMsg())
			return
		}
		group := rsp.GetGroup()
		gid, err := primitive.ObjectIDFromHex(group.GroupId)
		if err != nil {
			errCh <- err
			return
		}

		for _, clientID := range group.GetClientIDs() {
			// loop and send in kafka queue
			if clientID == event.FromID {
				continue
			}
			tid, err := primitive.ObjectIDFromHex(clientID)
			if err != nil {
				errCh <- err
				return
			}
			msg := models.Message{
				ID:        primitive.NewObjectID(),
				FromID:    cid,
				ToID:      tid,
				GroupID:   gid,
				Data:      event.Content,
				Timestamp: time.Now().Unix(),
			}
			b, err := json.Marshal(msg)
			if err != nil {
				errCh <- err
				return
			}
			msgCh <- kafka.Message{Value: b}
			// msg := kafka.Message{}
		}
		doneCh <- struct{}{}
	}()
	msgs := []kafka.Message{}
	for {
		select {
		case msg := <-msgCh:
			msgs = append(msgs, msg)
		case err := <-errCh:
			return err
		case <-doneCh:
			_, err := s.kafkaConn.WriteMessages(msgs...)
			if err != nil {
				return err
			}
			return nil
		}
	}
}

func (s *groupMessageServiceImpl) GetMessageFromGroup(groupID string, ClientID string, offset int64, limit int64) ([]*models.Message, error) {
	// Request to repository
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	req := &pb.GetGroupWithIDRequest{
		GroupId: groupID,
	}

	rsp, err := s.groupClient.GetGroupWithID(ctx, req)
	if err != nil {
		return nil, err
	}
	log.Println(rsp.GetGroup().String())

	if rsp.Error {
		return nil, fmt.Errorf("%v", rsp.GetMsg())
	}

	var (
		group = rsp.GetGroup()
		cids  = group.GetClientIDs()
	)
	for _, cid := range cids {
		if ClientID == cid {
			log.Println("Clientin group", ClientID)
			lastTimeStamp, err := s.groupMessageRepo.GetClientTimestamp(groupID, ClientID)
			if err != nil {
				log.Println("client not have timestamp")
			}
			log.Println("Update client timestamp")
			// TODO: check, msgs not found with timestamp
			if err := s.groupMessageRepo.UpdateClientTimestamp(groupID, ClientID, time.Now().Unix()); err != nil {
				log.Printf(err.Error())
			}
			return s.groupMessageRepo.GetMessageFromGroups(groupID, lastTimeStamp, offset, limit)
		}
	}
	return nil, fmt.Errorf("client %v not in group\n", ClientID)
}
