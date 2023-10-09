package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type GroupModel struct {
	GroupID   primitive.ObjectID `json:"group_id,omiempty" bson:"_id,omiempty"`
	Name      string             `json:"group_name,omiempty" bson:"group_name,omiempty"`
	IsPrivate bool               `json:"is_private,omiempty" bson:"is_private,omiempty"`
	ClientsID []string           `json:"clients,omiempty" bson:"clients,omiempty"`
	AdminsID  []string           `json:"admins,omiempty" bson:"admins,omiempty"`
}
type ClientModel struct {
	ClientID primitive.ObjectID `json:"client_id,omiempty" bson:"_id,omiempty"`
}
