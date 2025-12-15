package internal

import "github.com/gofiber/fiber/v2"

func IndexPage(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{})
}

func NotFoundPage(c *fiber.Ctx) error {
	return c.SendStatus(404)
}
