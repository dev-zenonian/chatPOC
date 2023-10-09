package transport

import (
	"common/transport"
	"common/utils"
	"log"
	"messageService/handler"
	"mock/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type messageHTTPTransportImpl struct {
	app            *fiber.App
	messageHandler handler.MessageHTTPHandler
}

func NewMessageHTTPTransportImpl(messageHandler handler.MessageHTTPHandler) (transport.HTTPTransport, error) {
	tsp := &messageHTTPTransportImpl{
		messageHandler: messageHandler,
	}
	tsp.app = utils.InitConnection()
	tsp.initRoute()
	return tsp, nil
}

func (s *messageHTTPTransportImpl) initRoute() {
	s.app.Use(cors.New())
	s.app.Use(logger.New())
	s.app.Post("/api/v1/message", middleware.MockMiddleware(), s.messageHandler.PostMessageHandler())
	s.app.Get("/api/v1/message", middleware.MockMiddleware(), s.messageHandler.GetMessageHandler())
}

func (s *messageHTTPTransportImpl) Listen(port string) error {
	log.Printf("HTTP Listening on port %v\n", port)
	return s.app.Listen(port)
}
