package handlers

import (
	"github.com/gofiber/fiber/v2"
	"zeroagencytest/pkg/repository"
	"zeroagencytest/pkg/utils/logging"
)

func ListHandler(ctx *fiber.Ctx, repo *repository.Repository) error {

	page := ctx.QueryInt("page", 1)
	limit := ctx.QueryInt("limit", 50)
	news, err := repo.Getlist(ctx.UserContext(), page, limit)
	if err != nil {
		logging.GetLogger().Errorln(err)
		return ctx.JSON(fiber.Map{
			"success": false,
			"error":   err,
		})
	}
	return ctx.JSON(fiber.Map{
		"Success": true,
		"News":    news,
	})
}
