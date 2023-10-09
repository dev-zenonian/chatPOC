package middleware

import (
	"common/models"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddUserMiddelware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// For testing
		log.Println("middleware")
		user := &models.ClientModel{
			ClientID: primitive.NewObjectID(),
		}
		log.Println("middleware of user", user)
		ctx.Locals("user", user)
		return ctx.Next()
	}
}

func MockMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		h := ctx.GetReqHeaders()
		auth, ok := h["Authorization"]
		if !ok {
			log.Println("Header section needed")
			return fiber.ErrForbidden
		}
		clientID := strings.Split(auth, " ")[1]
		cid, err := primitive.ObjectIDFromHex(clientID)
		if err != nil {
			log.Println("Middleware::", err)
			return fiber.ErrForbidden
		}
		log.Println("user: ", cid)
		user := &models.ClientModel{
			ClientID: cid,
		}
		ctx.Locals("user", user)
		return ctx.Next()
	}
}
