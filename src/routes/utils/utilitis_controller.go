package utils

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func oclock(c *fiber.Ctx) error {
	m := make(map[string]interface{})
	m["time"] = time.Now().Local()
	return c.JSON(m)
}
