package repository

import (
	"common/config"
	"common/models"
	"common/utils"

	"go.mongodb.org/mongo-driver/mongo"
)

type MessageRepository interface {
	GetInvidualMessage(groupID string, offset int64, limit int64) ([]*models.Message, error)
	SaveMessage(m *models.Message) error
	GetMessagesWithClientIDAndStatus(clientID string, status models.Status) ([]*models.Message, error)
	UpdateMessageStatus(mid string, status models.Status) error
}
type messageRepositoryImpl struct {
	db     *mongo.Database
	config *config.MongoConfig
}

func NewMessageRepositoryImpl(cfg *config.MongoConfig) (MessageRepository, error) {
	db, err := utils.InitMongoConnection(cfg.DSN, cfg.Database)
	if err != nil {
		return nil, err
	}
	repository := &messageRepositoryImpl{
		config: cfg,
		db:     db,
	}
	return repository, nil
}
