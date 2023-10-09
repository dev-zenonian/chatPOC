package transport

import (
	"common/models"
	"common/transport"
	"common/utils"
	"log"
	mockmiddleware "mock/middleware"
	"wsService/handler"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type httpTransportImpl struct {
	app       *fiber.App
	wsHandler handler.WSHTTPHandler
}

func NewHTTPTransportImpl(wsHandler handler.WSHTTPHandler) (transport.HTTPTransport, error) {
	tsp := &httpTransportImpl{
		wsHandler: wsHandler,
	}
	tsp.app = utils.InitConnection()
	tsp.initRoute()
	return tsp, nil
}

func (s *httpTransportImpl) initRoute() {
	s.app.Use(cors.New())
	s.app.Use(logger.New())
	s.app.Use("/ws", mockmiddleware.MockMiddleware(), func(ctx *fiber.Ctx) error {
		usr, err := utils.GetFromFiberLocal[models.ClientModel](ctx, "user")
		if err != nil {
			return fiber.ErrForbidden
		}
		log.Println("user: ", usr)
		if websocket.IsWebSocketUpgrade(ctx) {
			ctx.Locals("allowed", true)
			return ctx.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	s.app.Get("/ws/messenger", s.wsHandler.ServeWSS())
}

func (s *httpTransportImpl) Listen(port string) error {
	log.Printf("HTTP listening on %v\n", port)
	return s.app.Listen(port)
}
