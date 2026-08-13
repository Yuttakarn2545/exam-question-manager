package handler

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"exam-question-manager/internal/model"
	"exam-question-manager/internal/service"
)

type ExamHandler struct {
	service *service.ExamService
}

func NewExamHandler(service *service.ExamService) *ExamHandler {
	return &ExamHandler{service: service}
}

func (h *ExamHandler) Register(router fiber.Router) {
	router.Get("/exams", h.list)
	router.Post("/exams", h.create)
	router.Delete("/exams/:id", h.delete)
}

// list handles GET /exams?page=1&pageSize=10
func (h *ExamHandler) list(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	pageSize := c.QueryInt("pageSize", service.DefaultPageSize)

	return c.JSON(h.service.List(page, pageSize))
}

func (h *ExamHandler) create(c *fiber.Ctx) error {
	var input model.ExamInput
	if err := c.BodyParser(&input); err != nil {
		return respondError(c, fiber.StatusBadRequest, "INVALID_REQUEST", "invalid request body")
	}

	exam, err := h.service.Create(input)
	if err != nil {
		var verr *service.ValidationError
		if errors.As(err, &verr) {
			return respondError(c, fiber.StatusBadRequest, "VALIDATION_ERROR", verr.Message)
		}
		return respondError(c, fiber.StatusInternalServerError, "INTERNAL_ERROR", err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(exam)
}

func (h *ExamHandler) delete(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return respondError(c, fiber.StatusBadRequest, "INVALID_REQUEST", "invalid id")
	}

	if err := h.service.Delete(id); err != nil {
		if errors.Is(err, service.ErrExamNotFound) {
			return respondError(c, fiber.StatusNotFound, "NOT_FOUND", err.Error())
		}
		return respondError(c, fiber.StatusInternalServerError, "INTERNAL_ERROR", err.Error())
	}

	return c.SendStatus(fiber.StatusNoContent)
}
