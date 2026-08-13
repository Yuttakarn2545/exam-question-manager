package handler

import "github.com/gofiber/fiber/v2"

// respondError writes a consistent {"error": {"code", "message"}} body
// so the frontend can branch on `code` instead of parsing message strings.
func respondError(c *fiber.Ctx, status int, code, message string) error {
	return c.Status(status).JSON(fiber.Map{
		"error": fiber.Map{
			"code":    code,
			"message": message,
		},
	})
}
