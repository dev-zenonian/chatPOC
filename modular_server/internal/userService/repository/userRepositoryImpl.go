package repository

import (
	"common/models"
	"common/utils"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *userRepositoryImpl) GetUserWithID(userID string) (*models.UserModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	rid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": rid}
	usr, err := utils.GetOneObectWithFilter[models.UserModel](ctx, r.db.Collection("users"), filter)
	if err != nil {
		if err == mongo.ErrNilDocument {
			return nil, fmt.Errorf("User %v not found", userID)
		}
		return nil, err
	}
	return usr, nil
}

func (r *userRepositoryImpl) GetUsers() ([]*models.UserModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	usrs, err := utils.GetObjectsWithFilter[models.UserModel](ctx, r.db.Collection("users"), bson.M{})
	if err != nil {
		if err == mongo.ErrNilDocument {
			return []*models.UserModel{}, nil
		}
		return nil, err
	}
	return usrs, nil
}

func (r *userRepositoryImpl) InsertUser(user *models.UserModel) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if _, err := r.db.Collection("users").InsertOne(ctx, user); err != nil {
		return err
	}
	return nil
}
