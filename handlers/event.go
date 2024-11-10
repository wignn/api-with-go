package handlers

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/wignn/Native/models"
)

type EventHandler struct {
	repository models.EventRepository
}

func (h *EventHandler) GetMany(ctx *fiber.Ctx) error {
context, cencel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
defer cencel()

events, err := h.repository.GetMany(context)

if err != nil {
	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"message": "Internal Server Error",
	})
}

return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
	"status": "success",
	"message": "Get all events",
	"data":   events,
})
}

func (h *EventHandler) CreateOne(ctx *fiber.Ctx) error {
return nil
}
func (h *EventHandler) GetOne(ctx *fiber.Ctx) error {
return nil
}

func NewEventHandler(router fiber.Router, repository models.EventRepository) {
	handler := &EventHandler{
		repository: repository,
	}

	router.Get("/", handler.GetMany)
	router.Post("/", handler.CreateOne)
	router.Get("/:eventId", handler.GetOne)
}
