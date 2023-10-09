package models

import (
	"fmt"
	"strings"
)

type MsgType int64

const (
	InvidualMessage MsgType = 0
	GroupMessage    MsgType = 1
	JoinGroup       MsgType = 2
	LeaveGroup      MsgType = 3
)

func (t MsgType) String() string {
	switch t {
	case InvidualMessage:
		return "InvidualMessage"
	case GroupMessage:
		return "GroupMessage"
	case JoinGroup:
		return "JoinGroup"
	case LeaveGroup:
		return "LeaveGroup"
	default:
		return "Unknow"
	}
}

type Action struct {
	Type    MsgType `json:"type"`
	Content string  `json:"data"`
}

func (a Action) String() string {
	str := strings.Builder{}
	str.WriteString(fmt.Sprintf("Type: %s\n", a.Type))
	str.WriteString(fmt.Sprintf("Content: %v\n", a.Content))
	return str.String()
}

type Event struct {
	FromID    string `json:"from_id"`
	ToID      string `json:"to_id"`
	Action    `json:"action"`
	Timestamp int64 `json:"timestamp"`
}

func (e Event) String() string {
	str := strings.Builder{}
	str.WriteString(fmt.Sprintf("FromID: %v\n", e.FromID))
	str.WriteString(fmt.Sprintf("ToID: %v\n", e.ToID))
	str.WriteString(fmt.Sprintf("Action: %v\n", e.Action))
	return str.String()
}
