package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Message struct {
	ID      primitive.ObjectID `json:"message_id,omiempty" bson:"_id,omiempty"`
	FromID  primitive.ObjectID `json:"from_id,omiempty" bson:"fid,omiempty"`
	ToID    primitive.ObjectID `json:"to_id,omiempty" bson:"tid,omiempty"`
	Type    MsgType            `json:"message_type,omiempty" bson:"message_type,omiempty"`
	GroupID primitive.ObjectID `json:"group_id,omiempty" bson:"gid,omiempty"`
	// Status: "deliveried" -> "received" -> "seened"
	Status    Status `json:"status,omiempty" bson:"status,omiempty"`
	Data      string `json:"data,omiempty" bson:"data,omiempty"`
	Timestamp int64  `json:"timestamp,omiempty" bson:"timestamp,omiempty"`
}

type Status int64

const (
	Deliveried Status = 0
	Received   Status = 1
	// Seen       Status = 2
)

func (s Status) String() string {
	switch s {
	case Deliveried:
		return "deliveried"
	case Received:
		return "received"
	// case Seen:
	// 	return "seen"
	default:
		return "Unknown"
	}
}
