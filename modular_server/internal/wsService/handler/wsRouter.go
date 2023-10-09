package handler

import (
	"common/models"
	"common/pb"
	"context"
	"log"
	"wsService/service"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

// HTTP WSHTTPHandler for ws router
type WSHTTPHandler interface {
	ServeWSS() fiber.Handler
}

type wsHTTPHandlerImpl struct {
	wsService service.WSService
}

func NewWSHTTPHandlerImpl(wsService service.WSService) (WSHTTPHandler, error) {
	router := &wsHTTPHandlerImpl{
		wsService: wsService,
	}
	return router, nil
}

func (r *wsHTTPHandlerImpl) ServeWSS() fiber.Handler {
	return websocket.New(func(conn *websocket.Conn) {
		defer conn.Close()
		conn.SetCloseHandler(func(code int, text string) error {
			log.Printf("Received message from client: %v, code: %v\n", text, code)
			return conn.Close()
		})
		c := conn.Locals("user")
		client, ok := c.(*models.ClientModel)
		if !ok {
			log.Println("can't ustructure to clientmodel")
			return
		}

		log.Println("clientID", client.ClientID)
		// var (
		// 	clientID = primitive.NewObjectID().Hex()
		// )
		r.wsService.ServeClient(client.ClientID.Hex(), conn)
	})
}

type WSGPRCHandler struct {
	wsService service.WSService
	pb.WSHandlerServiceServer
}

func NewWSGRPCHandlerImpl(wsService service.WSService) (*WSGPRCHandler, error) {
	hdl := &WSGPRCHandler{
		wsService: wsService,
	}
	return hdl, nil
}

func (h *WSGPRCHandler) mustEmbedUnimplementedWSHandlerServiceServer() {}

func (h *WSGPRCHandler) PassMessageToClient(ctx context.Context, req *pb.PassMessageToClientRequest) (*pb.PassMessageToClientResponse, error) {
	var (
		fromID     = req.GetMessage().GetFromID()
		toID       = req.GetMessage().GetToID()
		msgContent = req.GetMessage().GetContent()
		groupID    = req.GetMessage().GetGroupID()
		msgID      = req.GetMessageID()
	)
	var err error
	log.Println("Received message to client::", req)
	if groupID != "" {
		err = h.wsService.ServeMessage(msgID, fromID, toID, msgContent, groupID)
	} else {
		err = h.wsService.ServeMessage(msgID, fromID, toID, msgContent)
	}
	if err != nil {
		rsp := &pb.PassMessageToClientResponse{
			Error: true,
			Data:  err.Error(),
		}
		return rsp, nil
	}
	rsp := &pb.PassMessageToClientResponse{
		Error: false,
		Data:  "",
	}
	return rsp, nil
}
