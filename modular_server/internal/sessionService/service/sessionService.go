package service

import (
	"fmt"
	"log"
	"sessionService/repository"
)

type SessionService interface {
	UpdateClientSession(clientID string, handlerID string, HandlerAddress string, Action string) error
	GetSessions() (*SessionInformation, error)
	GetsessionWithClientID(clientID string) (*SessionInformation, error)
}

type HandlerInformation struct {
	ClientIDs []string `json:"client_ids,omiempty"`
}
type SessionInformation struct {
	HandlersAddress map[string]HandlerInformation `json:"handlerID_clientIDs"`
}

type sessionServiceImpl struct {
	sessionRepo repository.SessionRepository
}

func NewSessionServiceImpl(sessionRepo repository.SessionRepository) (SessionService, error) {
	service := &sessionServiceImpl{
		sessionRepo: sessionRepo,
	}
	return service, nil
}
func (s *sessionServiceImpl) UpdateClientSession(clientID string, handlerID string, HandlerAddress string, Action string) error {
	log.Printf("Update client session, clientID: %v, handlerID: %v, handlerAddress: %v,Action: %v\n", clientID, handlerID, HandlerAddress, Action)
	switch {
	case Action == "register":
		return s.sessionRepo.SaveClientSession(clientID, handlerID, HandlerAddress)
	case Action == "unregister":
		return s.sessionRepo.DeleteClientSession(clientID, handlerID)
	default:
		return fmt.Errorf("Action: %v not found", Action)
	}
}

func (s *sessionServiceImpl) GetSessions() (*SessionInformation, error) {
	mapHandler_CLient, err := s.sessionRepo.GetCurrentSessionStatus()
	if err != nil {
		return nil, err
	}
	sessionEntry := map[string]HandlerInformation{}
	for handlerAddress, clientsID := range mapHandler_CLient {
		sessionEntry[handlerAddress] = HandlerInformation{
			ClientIDs: clientsID,
		}
	}
	return &SessionInformation{HandlersAddress: sessionEntry}, nil
}

func (s *sessionServiceImpl) GetsessionWithClientID(clientID string) (*SessionInformation, error) {
	mapHandler_client, err := s.sessionRepo.GetSessionWithClientID(clientID)
	if err != nil {
		return nil, err
	}
	sessionEntry := map[string]HandlerInformation{}
	for handlerAddress, clientsID := range mapHandler_client {
		sessionEntry[handlerAddress] = HandlerInformation{
			ClientIDs: clientsID,
		}
	}
	return &SessionInformation{HandlersAddress: sessionEntry}, nil

}
