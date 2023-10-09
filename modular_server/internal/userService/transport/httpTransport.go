package transport

import (
	"common/transport"
	"common/utils"
	"log"
	"userService/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type httpTransportImpl struct {
	userHandler handler.UserHTTPHandler
	app         *fiber.App
}

func NewHTTPTransportImpl(userHandler handler.UserHTTPHandler) (transport.HTTPTransport, error) {
	tsp := &httpTransportImpl{
		userHandler: userHandler,
	}
	tsp.app = utils.InitConnection()
	tsp.initRoute()
	return tsp, nil
}

func (h *httpTransportImpl) initConnection() {
}

func (s *httpTransportImpl) initRoute() {
	s.app.Use(cors.New())
	s.app.Use(logger.New())
	s.app.Get("/api/v1/user/:id", s.userHandler.GetUserWithIDHandler())
	s.app.Get("/api/v1/user", s.userHandler.GetUsersHandler())
	s.app.Post("/api/v1/user", s.userHandler.CreateUserHandler())
}

func (s *httpTransportImpl) Listen(port string) error {
	log.Printf("HTTP listening on %v\n", port)
	return s.app.Listen(port)
}
