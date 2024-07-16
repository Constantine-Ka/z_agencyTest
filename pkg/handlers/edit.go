package handlers

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"zeroagencytest/pkg/repository"
	"zeroagencytest/pkg/utils/logging"
)

func EditHandler(ctx *fiber.Ctx, repo *repository.Repository) error {
	logger := logging.GetLogger()
	id, err := ctx.ParamsInt("Id", 0)
	if err != nil || id == 0 {
		logger.Warnln("Invalid id parameter", err)
		return err
	}

	var body repository.News
	err = ctx.BodyParser(&body)
	if err != nil {
		logger.Errorln("Body Unmarshalling", err)
		return err
	}
	err = repo.UpdateElem(ctx.UserContext(), id, body)
	if err != nil {
		logger.Errorln("UpdateElem", err)
		return err
	}

	return ctx.SendStatus(http.StatusOK)
}
