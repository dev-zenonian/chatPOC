package repository

import (
	"common/config"
	"common/models"
	"common/utils"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	GetUserWithID(userID string) (*models.UserModel, error)
	GetUsers() ([]*models.UserModel, error)
	InsertUser(user *models.UserModel) error
}

type userRepositoryImpl struct {
	config *config.MongoConfig
	db     *mongo.Database
}

func NewUserRepositoryImpl(cfg *config.MongoConfig) (UserRepository, error) {
	db, err := utils.InitMongoConnection(cfg.DSN, cfg.Database)
	if err != nil {
		return nil, err
	}
	repo := &userRepositoryImpl{
		config: cfg,
		db:     db,
	}
	return repo, nil
}
