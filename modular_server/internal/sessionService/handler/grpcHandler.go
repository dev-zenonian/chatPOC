package handler

import (
	"common/pb"
	"context"
	"fmt"
	"log"
	"sessionService/service"
)

type SessionGRPCHandler struct {
	sessionService service.SessionService
	pb.UnimplementedSessionServiceServer
}

func NewSessionGRPCHandler(sessionService service.SessionService) (*SessionGRPCHandler, error) {
	sessionHandler := &SessionGRPCHandler{
		sessionService: sessionService,
	}
	return sessionHandler, nil
}
func (h *SessionGRPCHandler) RegisterClient(ctx context.Context, req *pb.RegisterClientRequest) (*pb.RegisterClientResponse, error) {
	var (
		clientID       = req.GetClientID()
		handlerID      = req.GetHandler().GetHandlerID()
		handlerAddress = req.GetHandler().GetHandlerAddress()
	)
	log.Printf("Register request clientID: %v, handlerID:%v, handlerAddress:%v\n", clientID, handlerID, handlerAddress)
	if clientID == "" || handlerAddress == "" || handlerID == "" {
		return &pb.RegisterClientResponse{
			Error: true,
			Data:  fmt.Sprintf("Request not valid, clientID = %v, handlerID = %v, handlerAddress = %v\n", clientID, handlerID, handlerAddress),
		}, nil
	}
	if err := h.sessionService.UpdateClientSession(clientID, handlerID, handlerAddress, "register"); err != nil {
		return &pb.RegisterClientResponse{
			Error: true,
			Data:  err.Error(),
		}, nil
	}
	return &pb.RegisterClientResponse{
		Error: false,
		Data:  "success",
	}, nil
}

func (h *SessionGRPCHandler) UnRegisterClient(ctx context.Context, req *pb.UnRegisterClientRequest) (*pb.UnRegisterClientResponse, error) {
	log.Printf("Unregister requst %v\n", req.String())
	var (
		clientID       = req.GetClientID()
		handlerID      = req.GetHandlerID()
		handlerAddress = req.GetHandlerAddress()
	)
	if clientID == "" || handlerAddress == "" || handlerID == "" {
		return &pb.UnRegisterClientResponse{
			Error: true,
			Data:  fmt.Sprintf("Request not valid, clientID = %v, handlerID = %v, handlerAddress = %v\n", clientID, handlerID, handlerAddress),
		}, nil
	}
	if err := h.sessionService.UpdateClientSession(clientID, handlerID, handlerAddress, "unregister"); err != nil {
		return &pb.UnRegisterClientResponse{
			Error: true,
			Data:  err.Error(),
		}, nil
	}
	return &pb.UnRegisterClientResponse{
		Error: false,
		Data:  "success",
	}, nil
}

func (h *SessionGRPCHandler) GetWSHandlerWithClientID(ctx context.Context, req *pb.GetWSHandlerWithClientIDsRequest) (*pb.GetWSHandlerWithClientIDsResponse, error) {
	var (
		clientID = req.GetClientID()
	)
	log.Println("Getsession forCLientiD::", clientID)
	ws, err := h.sessionService.GetsessionWithClientID(clientID)
	if err != nil {
		return &pb.GetWSHandlerWithClientIDsResponse{
			Error: true,
			Msg:   err.Error(),
		}, nil
	}
	log.Println("Distributor::GetWSHandlerWithClientID::", ws)
	handler := []*pb.HandlerInformation{}
	for handlerAddress, _ := range ws.HandlersAddress {
		handlerInformation := &pb.HandlerInformation{
			HandlerID:      handlerAddress,
			HandlerAddress: handlerAddress,
		}
		handler = append(handler, handlerInformation)

	}
	return &pb.GetWSHandlerWithClientIDsResponse{
		Error:    false,
		Handlers: handler,
	}, nil
}
