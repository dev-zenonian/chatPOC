package client

import (
	"sync"

	"github.com/gofiber/contrib/websocket"
)

type WebsocketClient interface {
	Accept(event any) error
	Id() string
	Conn() []*websocket.Conn
	AddConn(conn *websocket.Conn)
	RemoveConn(conn *websocket.Conn)
}

type websocketClientImpl struct {
	clientID string
	conn     map[*websocket.Conn]bool
	mu       sync.Mutex
}

func NewWebsocketClientImpl(clientID string, conn *websocket.Conn) WebsocketClient {
	client := &websocketClientImpl{
		clientID: clientID,
		conn:     map[*websocket.Conn]bool{conn: true},
		mu:       sync.Mutex{},
	}
	return client
}

func (c *websocketClientImpl) Accept(event any) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	for conn, onl := range c.conn {
		if !onl {
			continue
		}
		if err := conn.WriteJSON(event); err != nil {
			return err
		}
	}
	return nil
}

func (c *websocketClientImpl) Id() string {
	return c.clientID
}

func (c *websocketClientImpl) Conn() []*websocket.Conn {
	c.mu.Lock()
	defer c.mu.Unlock()
	conns := []*websocket.Conn{}
	for conn, onl := range c.conn {
		if !onl {
			continue
		}
		conns = append(conns, conn)
	}
	return conns
}

func (c *websocketClientImpl) AddConn(conn *websocket.Conn) {

	c.conn[conn] = true
}
func (c *websocketClientImpl) RemoveConn(conn *websocket.Conn) {
	delete(c.conn, conn)
}
