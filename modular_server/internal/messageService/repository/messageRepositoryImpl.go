package repository

import (
	"common/models"
	"common/utils"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *messageRepositoryImpl) GetInvidualMessage(groupID string, offset int64, limit int64) ([]*models.Message, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	rid, err := primitive.ObjectIDFromHex(groupID)
	if err != nil {
		return nil, err
	}
	filter := bson.M{
		"rid": rid,
	}
	opts := options.Find().SetSkip(offset).SetLimit(limit)
	message, err := utils.GetObjectsWithFilter[models.Message](ctx, r.db.Collection("messages"), filter, opts)
	return message, err
}

func (r *messageRepositoryImpl) SaveMessage(m *models.Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	_, err := r.db.Collection("messages").InsertOne(ctx, m)
	if err != nil {
		return err
	}
	return nil
}

func (r *messageRepositoryImpl) GetMessagesWithClientIDAndStatus(clientID string, status models.Status) ([]*models.Message, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	cid, err := primitive.ObjectIDFromHex(clientID)
	if err != nil {
		return nil, err
	}
	filter := bson.M{
		"tid":    cid,
		"status": status,
	}
	return utils.GetObjectsWithFilter[models.Message](ctx, r.db.Collection("messages"), filter)
}

func (r *messageRepositoryImpl) UpdateMessageStatus(messageID string, status models.Status) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	mid, err := primitive.ObjectIDFromHex(messageID)
	if err != nil {
		return err
	}
	filter := bson.M{
		"_id": mid,
	}
	update := bson.M{
		"$set": bson.M{
			"status": status,
		},
	}
	_, err = r.db.Collection("messages").UpdateOne(ctx, filter, update)
	return err
}
