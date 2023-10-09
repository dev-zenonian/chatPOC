package transport

import (
	"common/transport"
	"common/utils"
	"groupService/handler"
	"log"
	"mock/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type httpTransportImpl struct {
	app          *fiber.App
	groupHandler handler.GroupHTTPHandler
}

func NewHTTPTransportImpl(groupMessageHandler handler.GroupHTTPHandler) (transport.HTTPTransport, error) {
	tsp := &httpTransportImpl{
		groupHandler: groupMessageHandler,
	}
	tsp.app = utils.InitConnection()
	tsp.initRoute()
	return tsp, nil
}

func (s *httpTransportImpl) initRoute() {
	s.app.Use(cors.New())
	s.app.Use(logger.New())

	s.app.Post("/api/v1/group", middleware.MockMiddleware(), s.groupHandler.CreateGroupHandler())
	s.app.Delete("/api/v1/group", middleware.MockMiddleware(), s.groupHandler.DeleteGroupHandler())
	s.app.Get("/api/v1/group", s.groupHandler.GetGroupsHandler())
	s.app.Get("/api/v1/group/:id", s.groupHandler.GetGroupWithIDHandler())
	s.app.Get("/api/v1/group/client", s.groupHandler.GetClientsOfGroupHandler())

	s.app.Post("/api/v1/group/register", middleware.MockMiddleware(), s.groupHandler.RegisterClientToGroup())
	s.app.Post("/api/v1/group/unregister", middleware.MockMiddleware(), s.groupHandler.UnRegisterClientFromGroup())

	s.app.Post("/api/v1/group/client", middleware.MockMiddleware(), s.groupHandler.AdClientsToGroupHandler())
	s.app.Delete("/api/v1/group/client", middleware.MockMiddleware(), s.groupHandler.RemovelientsToGroupHandler())
}

func (s *httpTransportImpl) Listen(port string) error {
	log.Printf("HTTP Listening on port %v\n", port)
	return s.app.Listen(port)
}
