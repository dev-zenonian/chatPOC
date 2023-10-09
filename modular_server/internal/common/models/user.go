package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserModel struct {
	UserID   primitive.ObjectID `json:"user_id,omiempty" bson:"_id,omiempty"`
	UserName string             `json:"user_name,omiempty" bson:"user_name,omiempty"`
}
