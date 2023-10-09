package transport

import (
	"common/transport"
	"common/utils"
	"groupMessageService/handler"
	"log"
	"mock/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type httpTransportImpl struct {
	app                 *fiber.App
	groupMessageHandler handler.GroupMessageHandler
}

func NewHTTPTransportImpl(groupMessageHandler handler.GroupMessageHandler) (transport.HTTPTransport, error) {
	tsp := &httpTransportImpl{
		groupMessageHandler: groupMessageHandler,
	}
	tsp.app = utils.InitConnection()
	tsp.initRoute()
	return tsp, nil
}

func (s *httpTransportImpl) initRoute() {
	s.app.Use(cors.New())
	s.app.Use(logger.New())

	s.app.Post("/api/v1/groupmessage", middleware.MockMiddleware(), s.groupMessageHandler.PostGroupMessageHandler())
	s.app.Get("/api/v1/groupmessage", middleware.MockMiddleware(), s.groupMessageHandler.GetGroupMessageHandler())
}

func (s *httpTransportImpl) Listen(port string) error {
	log.Printf("Listening on port: %v", port)
	return s.app.Listen(port)
}
