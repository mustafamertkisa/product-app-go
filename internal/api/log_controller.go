package api

import (
	"product-app-go/internal/application/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LogController struct {
	logService service.LogService
}

func NewLogController(service service.LogService) *LogController {
	return &LogController{logService: service}
}

func (controller *LogController) GetLogById(ctx *fiber.Ctx) error {
	logIdParam := ctx.Params("id")
	logId, err := primitive.ObjectIDFromHex(logIdParam)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	log, err := controller.logService.GetLogByLogId(logId)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(log)
}

func (controller *LogController) GetAllLogs(ctx *fiber.Ctx) error {
	logs, err := controller.logService.GetAllLogs()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(logs)
}

func (controller *LogController) GetLogsByUserId(ctx *fiber.Ctx) error {
	userIdParam := ctx.Params("userId")
	userId, err := strconv.Atoi(userIdParam)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid user ID",
		})
	}

	logs, err := controller.logService.GetLogsByUserId(userId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(logs)
}

func (controller *LogController) DeleteLogById(ctx *fiber.Ctx) error {
	logIdParam := ctx.Params("id")
	logId, err := primitive.ObjectIDFromHex(logIdParam)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err = controller.logService.DeleteLogById(logId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Log deleted successfully",
	})
}

func (controller *LogController) DeleteAllLogs(ctx *fiber.Ctx) error {
	err := controller.logService.DeleteAllLogs()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "All logs deleted successfully",
	})
}
