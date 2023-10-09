package utils

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetFromFiberLocal[T any](ctx *fiber.Ctx, key string, value ...interface{}) (*T, error) {
	ele := ctx.Locals(key, value...)
	res, ok := ele.(*T)
	if !ok {
		return nil, fmt.Errorf("Cann't destructor to interface")
	}
	return res, nil
}

func InitConnection() *fiber.App {
	config := fiber.Config{
		ReadTimeout:           5 * time.Second,
		WriteTimeout:          5 * time.Second,
		ReadBufferSize:        2e6,
		WriteBufferSize:       2e6,
		DisableStartupMessage: true,
	}
	app := fiber.New(config)
	return app
}
