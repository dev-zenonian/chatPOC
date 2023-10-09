package service

import (
	"common/models"
	"userService/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService interface {
	GetUserWithID(userID string) (*models.UserModel, error)
	GetUsers() ([]*models.UserModel, error)
	CreateUser(UserName string) (*models.UserModel, error)
}

type userServiceImpl struct {
	repo repository.UserRepository
}

func NewUserServiceImpl(userRepo repository.UserRepository) (UserService, error) {
	service := &userServiceImpl{
		repo: userRepo,
	}
	return service, nil
}

func (s *userServiceImpl) GetUserWithID(userID string) (*models.UserModel, error) {
	return s.repo.GetUserWithID(userID)
}
func (s *userServiceImpl) GetUsers() ([]*models.UserModel, error) {
	return s.repo.GetUsers()
}

func (s *userServiceImpl) CreateUser(userName string) (*models.UserModel, error) {
	usr := &models.UserModel{
		UserID:   primitive.NewObjectID(),
		UserName: userName,
	}
	if err := s.repo.InsertUser(usr); err != nil {
		return nil, err
	}
	return usr, nil
}
