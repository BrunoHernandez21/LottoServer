package utils

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func oclock(c *fiber.Ctx) error {
	m := make(map[string]string)
	m["time"] = time.Now().Format("15:04:05.00")
	return c.JSON(m)
}
