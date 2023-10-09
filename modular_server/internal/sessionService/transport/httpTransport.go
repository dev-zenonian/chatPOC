package transport

import (
	"common/transport"
	"common/utils"
	"log"
	"sessionService/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type httpTransportImpl struct {
	app            *fiber.App
	sessionHandler handler.SessionHTTPHandler
}

func NewHTTPTransportImpl(sessionHandler handler.SessionHTTPHandler) (transport.HTTPTransport, error) {
	tsp := &httpTransportImpl{
		sessionHandler: sessionHandler,
	}
	tsp.app = utils.InitConnection()
	tsp.initRoute()
	return tsp, nil
}

func (s *httpTransportImpl) initRoute() {
	s.app.Use(cors.New())
	s.app.Use(logger.New())
	s.app.Get("/api/v1/session", s.sessionHandler.GetSessionHandler())
	s.app.Get("/api/v1/session/:clientID", s.sessionHandler.GetSessionWithClientIDHandler())
}

func (s *httpTransportImpl) Listen(port string) error {
	log.Printf("HTTP Listening on port %v\n", port)
	return s.app.Listen(port)
}
