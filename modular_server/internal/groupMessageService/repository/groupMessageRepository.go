package repository

import (
	"common/config"
	"common/models"
	"common/utils"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GroupMessageRepository interface {
	GetMessageFromGroups(groupID string, timestamp, offset, limit int64) ([]*models.Message, error)
	SaveMessageInGroups(message *models.Message) error
	SaveClientTimestamp(groupID string, clientID string, timestamp int64) error
	GetClientTimestamp(groupID string, clientID string) (int64, error)
	UpdateClientTimestamp(groupID string, clientID string, timeStamp int64) error
}

type groupMessageRepositoryImpl struct {
	db     *mongo.Database
	config *config.MongoConfig
}

func NewGroupMessageRepositoryImpl(cfg *config.MongoConfig) (GroupMessageRepository, error) {
	db, err := utils.InitMongoConnection(cfg.DSN, cfg.Database)
	if err != nil {
		return nil, err
	}
	repo := &groupMessageRepositoryImpl{
		config: cfg,
		db:     db,
	}
	return repo, nil
}

func (r *groupMessageRepositoryImpl) GetMessageFromGroups(groupID string, timestamp, offset, limit int64) ([]*models.Message, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	gid, err := primitive.ObjectIDFromHex(groupID)
	if err != nil {
		return nil, err
	}
	filter := bson.M{
		"gid":       bson.M{"$eq": gid},
		"timestamp": bson.M{"$gt": timestamp},
	}
	msgs, err := utils.GetObjectsWithFilter[models.Message](ctx, r.db.Collection("groupmessages"), filter, options.Find().SetLimit(limit).SetSkip(offset))
	return msgs, err
}
func (r *groupMessageRepositoryImpl) SaveMessageInGroups(message *models.Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	_, err := r.db.Collection("groupmessages").InsertOne(ctx, message)
	return err
}

type timeStampDocument struct {
	ClientID  primitive.ObjectID `bson:"client_id"`
	GroupID   primitive.ObjectID `bson:"group_id"`
	TimeStamp int64              `bson:"timestamp"`
}

func (r *groupMessageRepositoryImpl) SaveClientTimestamp(groupID string, clientID string, timestamp int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	gid, err := primitive.ObjectIDFromHex(groupID)
	if err != nil {
		return err
	}
	cid, err := primitive.ObjectIDFromHex(clientID)
	if err != nil {
		return err
	}
	doc := timeStampDocument{
		ClientID:  cid,
		GroupID:   gid,
		TimeStamp: timestamp,
	}
	_, err = r.db.Collection("grouptimestamps").InsertOne(ctx, doc)
	return err
}
func (r *groupMessageRepositoryImpl) UpdateClientTimestamp(groupID string, clientID string, timeStamp int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	gid, err := primitive.ObjectIDFromHex(groupID)
	if err != nil {
		return err
	}
	cid, err := primitive.ObjectIDFromHex(clientID)
	if err != nil {
		return err
	}
	filter := bson.M{
		"group_id":  bson.M{"$eq": gid},
		"client_id": bson.M{"$eq": cid},
	}
	update := bson.M{
		"$set": bson.M{"timestamp": timeStamp},
	}
	err = r.db.Collection("grouptimestamps").FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetUpsert(true)).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *groupMessageRepositoryImpl) GetClientTimestamp(groupID string, clientID string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	gid, err := primitive.ObjectIDFromHex(groupID)
	if err != nil {
		return 0, err
	}
	cid, err := primitive.ObjectIDFromHex(clientID)
	if err != nil {
		return 0, err
	}
	filter := bson.M{
		"group_id":  bson.M{"$eq": gid},
		"client_id": bson.M{"$eq": cid},
	}
	doc, err := utils.GetOneObectWithFilter[timeStampDocument](ctx, r.db.Collection("grouptimestamps"), filter)
	if err != nil {
		return 0, err
	}
	return doc.TimeStamp, nil
}
