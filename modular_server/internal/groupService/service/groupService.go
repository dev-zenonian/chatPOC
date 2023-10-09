package service

import (
	"common/models"
	"common/pb"
	"context"
	"fmt"
	"groupService/repository"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GroupService interface {
	CreateGroup(groupName string, AdminID string, clientID []string, isPrivate bool) (*models.GroupModel, error)
	DeleteGroup(groupId string, adminID string) error
	GetClientsOfGroup(groupID string) ([]*models.ClientModel, error)
	GetGroupWithID(groupID string) (*models.GroupModel, error)
	GetGroups() ([]*models.GroupModel, error)
	RemoveClientFromGroup(groupID string, clientID string) error
	AddClientToGroup(groupID string, clientID string) error
	AddClientsToGroupWithAdminID(groupId string, adminID string, clientsID []string) ([]string, error)
	RemoveClientsFromGroupWithAdminID(groupID string, adminID string, clientsID []string) ([]string, error)
}

type groupServiceImpl struct {
	groupRepo  repository.GroupRepository
	userClient pb.UserServiceClient
}

func NewGroupServiceImpl(groupRepository repository.GroupRepository, userClient pb.UserServiceClient) (GroupService, error) {
	service := &groupServiceImpl{
		groupRepo:  groupRepository,
		userClient: userClient,
	}
	return service, nil
}

func (s *groupServiceImpl) GetClientsOfGroup(groupId string) ([]*models.ClientModel, error) {
	group, err := s.groupRepo.GetGroupWithID(groupId)
	if err != nil {
		return nil, err
	}
	clients := []*models.ClientModel{}
	for _, client := range group.ClientsID {
		cid, err := primitive.ObjectIDFromHex(client)
		if err != nil {
			return nil, err
		}
		clients = append(clients, &models.ClientModel{ClientID: cid})
	}
	return clients, nil
}

func (s *groupServiceImpl) GetGroupWithID(groupID string) (*models.GroupModel, error) {
	return s.groupRepo.GetGroupWithID(groupID)
}
func (s *groupServiceImpl) GetGroups() ([]*models.GroupModel, error) {
	return s.groupRepo.GetGroups()
}
func (s *groupServiceImpl) CreateGroup(groupName string, AdminID string, clientIDs []string, isPrivate bool) (*models.GroupModel, error) {

	validatedClient, unvalidatedsClient, err := s.ValidateClients(clientIDs)
	if err != nil {
		log.Fatal(err)
	}
	if len(unvalidatedsClient) != 0 {
		log.Printf("clients %v unvalidate\n", unvalidatedsClient)
	}
	group := &models.GroupModel{
		GroupID:   primitive.NewObjectID(),
		Name:      groupName,
		IsPrivate: isPrivate,
		AdminsID:  []string{AdminID},
		ClientsID: validatedClient,
	}

	if err := s.groupRepo.CreateGroup(group); err != nil {
		return nil, err
	}
	return group, nil
}

func (s *groupServiceImpl) RemoveClientFromGroup(groupID string, clientID string) error {
	group, err := s.groupRepo.GetGroupWithID(groupID)
	if err != nil {
		return err
	}
	for _, cid := range group.ClientsID {
		if cid == clientID {
			if err := s.groupRepo.RemoveClientsFromGroup(groupID, []string{clientID}); err != nil {
				return err
			}
			return nil
		}
	}
	return fmt.Errorf("Client %v not in group %v\n", clientID, groupID)
}
func (s *groupServiceImpl) AddClientToGroup(groupID string, clientID string) error {
	group, err := s.groupRepo.GetGroupWithID(groupID)
	if err != nil {
		return err
	}
	validatedClient, _, err := s.ValidateClients([]string{clientID})
	if err != nil {
		return err
	}
	// check if clientID not validated
	if len(validatedClient) != 1 {
		return fmt.Errorf("ClientIDs %v cann't validate", clientID)
	}

	for _, client := range group.ClientsID {
		if client == clientID {
			//client already in group
			return fmt.Errorf("Client %v already in group %v\n", clientID, groupID)
		}
	}
	return s.groupRepo.AddClientsToGroup(groupID, []string{clientID})
}

// RPC relate

func (s *groupServiceImpl) AddClientsToGroupWithAdminID(groupId string, adminID string, clientsID []string) ([]string, error) {
	group, err := s.groupRepo.GetGroupWithID(groupId)
	if err != nil {
		return []string{}, err
	}
	var (
		adminsID = group.AdminsID
	)
	validatedClients, unvalidatedClients, err := s.ValidateClients(clientsID)
	if err != nil {
		return unvalidatedClients, err
	}
	if len(unvalidatedClients) != 0 {
		log.Printf("Clients %v unvalidated\n", unvalidatedClients)
	}
	err = fmt.Errorf("Admin %v doesn't have permission", adminID)
	for _, aid := range adminsID {
		if aid == adminID {
			err = nil
		}
	}
	if err != nil {
		return unvalidatedClients, err
	}

	err = s.groupRepo.AddClientsToGroup(groupId, validatedClients)
	return unvalidatedClients, err
}
func (s *groupServiceImpl) RemoveClientsFromGroupWithAdminID(groupID string, adminID string, clientsID []string) ([]string, error) {
	group, err := s.groupRepo.GetGroupWithID(groupID)
	if err != nil {
		return []string{}, err
	}
	var (
		adminsID = group.AdminsID
	)
	validatedClients, unValidatedClients, err := s.ValidateClients(clientsID)
	if err != nil {
		return unValidatedClients, err
	}
	if len(unValidatedClients) != 0 {
		log.Printf("clients %v unvalidate\n", unValidatedClients)
	}
	if err != nil {
		log.Println(err)
	}
	err = fmt.Errorf("Admin %v doesn't have permission", adminID)
	for _, aid := range adminsID {
		if aid == adminID {
			err = nil
		}
	}
	if err != nil {
		return unValidatedClients, err
	}

	err = s.groupRepo.RemoveClientsFromGroup(groupID, validatedClients)
	return unValidatedClients, err
}

func (s *groupServiceImpl) ValidateClients(clientsID []string) (validatedClients []string, unValidatedClients []string, err error) {
	validatedClients = []string{}
	unValidatedClients = []string{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	for _, clientId := range clientsID {
		req := &pb.GetUserWithIDRequest{
			UserId: clientId,
		}
		rsp, e := s.userClient.GetUserWithID(ctx, req)
		if e != nil {
			err = e
		}
		if rsp.Error || e != nil {
			unValidatedClients = append(unValidatedClients, clientId)
		} else {
			validatedClients = append(validatedClients, clientId)
		}
	}
	return
}

func (s *groupServiceImpl) DeleteGroup(groupId string, adminID string) error {
	group, err := s.groupRepo.GetGroupWithID(groupId)
	if err != nil {
		return err
	}
	err = fmt.Errorf("Admin %v not have permission\n", adminID)
	for _, aid := range group.AdminsID {
		if aid == adminID {
			err = nil
		}
	}

	if err != nil {
		return err
	}
	return s.groupRepo.DeleteGroup(groupId)
}
