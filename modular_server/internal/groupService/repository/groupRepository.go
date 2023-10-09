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
)

type GroupRepository interface {
	CreateGroup(group *models.GroupModel) error
	DeleteGroup(gid string) error
	GetGroupWithID(groupID string) (*models.GroupModel, error)
	GetGroups() ([]*models.GroupModel, error)
	RemoveClientsFromGroup(groupID string, clientsID []string) error
	AddClientsToGroup(groupID string, clientsID []string) error
}

type groupRepositoryImpl struct {
	config *config.MongoConfig
	db     *mongo.Database
}

func NewGroupRepositoryImpl(cfg *config.MongoConfig) (GroupRepository, error) {
	db, err := utils.InitMongoConnection(cfg.DSN, cfg.Database)
	if err != nil {
		return nil, err
	}
	repo := &groupRepositoryImpl{
		config: cfg,
		db:     db,
	}
	return repo, nil
}

func (r *groupRepositoryImpl) CreateGroup(group *models.GroupModel) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	_, err := r.db.Collection("groups").InsertOne(ctx, group)
	return err
}
func (r *groupRepositoryImpl) DeleteGroup(groupId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	gid, err := primitive.ObjectIDFromHex(groupId)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": gid}
	_, err = r.db.Collection("groups").DeleteOne(ctx, filter)
	return err
}

func (r *groupRepositoryImpl) GetGroupWithID(groupID string) (*models.GroupModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	rid, err := primitive.ObjectIDFromHex(groupID)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": rid}
	return utils.GetOneObectWithFilter[models.GroupModel](ctx, r.db.Collection("groups"), filter)
}

func (r *groupRepositoryImpl) GetGroups() ([]*models.GroupModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	return utils.GetObjectsWithFilter[models.GroupModel](ctx, r.db.Collection("groups"), bson.M{})
}

func (r *groupRepositoryImpl) RemoveClientsFromGroup(groupID string, clientsID []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	gid, err := primitive.ObjectIDFromHex(groupID)
	if err != nil {
		return err
	}
	cids := []primitive.ObjectID{}
	for _, clientID := range clientsID {
		cid, err := primitive.ObjectIDFromHex(clientID)
		if err != nil {
			return err
		}
		cids = append(cids, cid)
	}

	filter := bson.M{"_id": gid}
	update := bson.M{"$pullAll": bson.M{"clients": cids}}

	res := r.db.Collection("groups").FindOneAndUpdate(ctx, filter, update)
	if err := res.Err(); err != nil {
		return err
	}
	return nil
}

func (r *groupRepositoryImpl) AddClientsToGroup(groupID string, clientsID []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	gid, err := primitive.ObjectIDFromHex(groupID)
	if err != nil {
		return err
	}
	cids := []primitive.ObjectID{}
	for _, clientID := range clientsID {
		cid, err := primitive.ObjectIDFromHex(clientID)
		if err != nil {
			return err
		}
		cids = append(cids, cid)
	}

	filter := bson.M{"_id": gid}
	update := bson.M{"$addToSet": bson.M{"clients": bson.M{"$each": cids}}}

	err = r.db.Collection("groups").FindOneAndUpdate(ctx, filter, update).Err()
	return err
}
