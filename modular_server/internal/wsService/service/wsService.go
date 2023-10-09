package service

import (
	"common/models"
	"common/pb"
	"context"
	"fmt"
	"log"
	"sync"
	"time"
	"wsService/client"

	"github.com/gofiber/contrib/websocket"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WSService interface {
	ServeClient(clientID string, conn *websocket.Conn)
	ServeMessage(msgID string, fromID string, toID string, message string, groupID ...string) error
}

type clientRequest struct {
	ClientID       string          `json:"client_id,omiempty"`
	GroupID        string          `json:"group_id,omiempty"`
	HandlerID      string          `json:"handler_id,omiempty"`
	HandlerAddress string          `json:"handler_address,omiempty"`
	ClientConn     *websocket.Conn `json:"client_conn"`
}

type reqEvent struct {
	Action string `json:"action,omiempty"`
	Data   string `json:"data,omiempty"`
}

type wsServiceImpl struct {
	mu             sync.Mutex
	clients        map[string]client.WebsocketClient
	registerCh     chan *clientRequest
	unRegisterCh   chan *clientRequest
	eventCh        chan *models.Event
	messageCh      chan *models.Message
	sessionClient  pb.SessionServiceClient
	messageClient  pb.MessageServiceClient
	handlerID      string
	handlerAddress string
	kafkaConn      *kafka.Conn
}

func NewWSServiceImpl(
	sessionClient pb.SessionServiceClient,
	messageClient pb.MessageServiceClient,
	kafkaConn ...*kafka.Conn,
) (WSService, error) {

	wsService := &wsServiceImpl{
		clients:        map[string]client.WebsocketClient{},
		registerCh:     make(chan *clientRequest, 10),
		unRegisterCh:   make(chan *clientRequest, 10),
		eventCh:        make(chan *models.Event, 10),
		messageCh:      make(chan *models.Message, 10),
		mu:             sync.Mutex{},
		handlerID:      "sample",
		sessionClient:  sessionClient,
		messageClient:  messageClient,
		handlerAddress: "127.0.0.1:8083",
	}
	if len(kafkaConn) == 1 {
		wsService.kafkaConn = kafkaConn[0]
	}
	go wsService.BackgroundWorker()
	return wsService, nil

}

func (s *wsServiceImpl) BackgroundWorker() {
	for {
		select {
		case r := <-s.registerCh:
			{
				// send information to session Service
				ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
				defer cancel()
				sessionRequest := &pb.RegisterClientRequest{
					ClientID: r.ClientID,
					Handler: &pb.HandlerInformation{
						HandlerID:      s.handlerID,
						HandlerAddress: s.handlerAddress,
					},
				}
				rsp, err := s.sessionClient.RegisterClient(ctx, sessionRequest)
				if err != nil {
					log.Println(err.Error())
					r.ClientConn.Close()
					continue
				}
				if rsp.Error {
					log.Println(rsp.Data)
					r.ClientConn.Close()
					continue
				}
				go s.RegisterConn(r)
			}
		case r := <-s.unRegisterCh:
			{
				// send information to session Service
				ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
				defer cancel()
				sessionRequest := &pb.UnRegisterClientRequest{
					ClientID:       r.ClientID,
					HandlerID:      s.handlerID,
					HandlerAddress: s.handlerAddress,
				}
				rsp, err := s.sessionClient.UnRegisterClient(ctx, sessionRequest)
				if err != nil {
					r.ClientConn.Close()
					continue
				}
				if rsp.Error {
					r.ClientConn.Close()
					continue
				}
				go s.UnRegisterConn(r)
			}
		case e := <-s.eventCh:
			{
				continue
				log.Println(e)
				// Send event to message Service
				// BroadCast to ToID

				// b, err := json.Marshal(e)
				// if err != nil {
				// 	continue
				// }
				// if _, err := s.kafkaConn.WriteMessages(kafka.Message{Value: b}); err != nil {
				// 	log.Println(err)
				// 	continue
				// }
				// log.Println("BroadCast event", e)
			}
		case msg := <-s.messageCh:
			{
				// received message to send to client
				s.SendMessage(msg)
			}
		}
	}

}

func (s *wsServiceImpl) RegisterConn(req *clientRequest) {
	log.Printf("Register for client: %v ,to group: %v", req.ClientID, req.GroupID)
	s.mu.Lock()
	c, ok := s.clients[req.ClientID]
	if !ok {
		s.clients[req.ClientID] = client.NewWebsocketClientImpl(req.ClientID, req.ClientConn)

		s.mu.Unlock()
		return
	}
	go c.AddConn(req.ClientConn)
	s.mu.Unlock()
}

func (s *wsServiceImpl) UnRegisterConn(req *clientRequest) {

	log.Printf("Unregister for client: %v ,from group: %v", req.ClientID, req.GroupID)
	s.mu.Lock()
	c, ok := s.clients[req.ClientID]
	if !ok {
		//client not exits
		s.mu.Unlock()
		log.Println("Unregister but client not exists")
		return
	}
	go c.RemoveConn(req.ClientConn)
	s.mu.Unlock()
}

func (s *wsServiceImpl) ServeMessage(messageID string, fromID string, toID string, message string, groupID ...string) error {
	fid, err := primitive.ObjectIDFromHex(fromID)
	if err != nil {
		return err
	}
	msgid, err := primitive.ObjectIDFromHex(messageID)
	if err != nil {
		return err
	}
	tid, err := primitive.ObjectIDFromHex(toID)
	if err != nil {
		return err
	}

	gid := primitive.NilObjectID
	if len(groupID) == 1 {
		gid, err = primitive.ObjectIDFromHex(groupID[0])
		if err != nil {
			return err
		}
	}

	msg := &models.Message{
		ID:        msgid,
		FromID:    fid,
		ToID:      tid,
		GroupID:   gid,
		Data:      message,
		Timestamp: time.Now().Unix(),
	}
	s.messageCh <- msg
	return nil
}

func (s *wsServiceImpl) SendMessage(msg *models.Message) error {
	client, ok := s.clients[msg.ToID.Hex()]
	if !ok {
		log.Println("there no client with clientid", msg.ToID.Hex())
		return nil
	}
	log.Println("received message")
	return client.Accept(msg)
}

func (s *wsServiceImpl) GetUnreadMessage(clientID string) ([]*models.Message, error) {
	req := &pb.GetMessageWithStatusRequest{
		ClientID: clientID,
		Status:   pb.Status_Deliveried,
	}
	rsp, err := s.messageClient.GetMessageWithStatus(context.TODO(), req)
	if err != nil {
		log.Println("WSService::GetMessageWithStatus::RSP", err.Error())
		return nil, err
	}
	if rsp.Error {
		log.Println("WSService::GetMessageWithStatus::RSP", err.Error())
		return nil, fmt.Errorf(rsp.GetMsg())
	}
	protoMsgs := rsp.GetMessages()
	res := []*models.Message{}
	for _, protoMsg := range protoMsgs {
		mid, err := primitive.ObjectIDFromHex(protoMsg.GetMessageID())
		if err != nil {
			return nil, err
		}
		fid, err := primitive.ObjectIDFromHex(protoMsg.GetFromID())
		if err != nil {
			return nil, err
		}
		tid, err := primitive.ObjectIDFromHex(protoMsg.GetToID())
		if err != nil {
			return nil, err
		}
		gid, err := primitive.ObjectIDFromHex(protoMsg.GetGroupID())
		if err != nil {
			gid = primitive.NilObjectID
		}
		msg := &models.Message{
			ID:        mid,
			FromID:    fid,
			ToID:      tid,
			GroupID:   gid,
			Data:      protoMsg.GetContent(),
			Timestamp: protoMsg.GetTimestamp(),
		}
		res = append(res, msg)
	}
	return res, nil
}

func (s *wsServiceImpl) ServeClient(clientID string, conn *websocket.Conn) {
	var (
		req = &clientRequest{
			ClientID:   clientID,
			ClientConn: conn,
		}
	)
	go func() {
		s.registerCh <- req
		event := &models.Event{
			FromID:    req.ClientID,
			ToID:      req.GroupID,
			Timestamp: time.Now().Unix(),
			Action: models.Action{
				Type:    models.JoinGroup,
				Content: fmt.Sprintf("member: %v, joined", req.ClientID[len(req.ClientID)-5:]),
			},
		}
		s.eventCh <- event

	}()
	go func() {
		msgs, err := s.GetUnreadMessage(clientID)
		if err != nil {
			log.Print("WsHandler::ServeClient::GetUnreadMessage::", err.Error())
		}
		if len(msgs) > 0 {
			log.Println("there are unread message")
			for _, msg := range msgs {
				s.messageCh <- msg
			}
		}
	}()

	defer func() {
		s.unRegisterCh <- req
		event := &models.Event{
			FromID:    req.ClientID,
			ToID:      req.GroupID,
			Timestamp: time.Now().Unix(),
			Action: models.Action{
				Type:    models.LeaveGroup,
				Content: fmt.Sprintf("member: %v, joined", req.ClientID[len(req.ClientID)-5:]),
			},
		}
		s.eventCh <- event
	}()

	for {
		var req = &reqEvent{}
		if err := conn.ReadJSON(req); err != nil {
			return
		}

		t := models.InvidualMessage
		if req.Action == "group" {
			t = models.GroupMessage
		}
		event := &models.Event{
			FromID: clientID,
			Action: models.Action{
				Type:    t,
				Content: fmt.Sprintf("%v: %v", clientID[len(clientID)-5:], req.Data),
			},
			Timestamp: time.Now().Unix(),
		}
		s.eventCh <- event
	}
}
