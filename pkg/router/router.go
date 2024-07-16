package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"zeroagencytest/pkg/handlers"
	"zeroagencytest/pkg/repository"
	"zeroagencytest/pkg/utils/logging"
)

func Init(db *repository.Repository, port string) {
	logger := logging.GetLogger()

	app := fiber.New()
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"user":  "123456",
			"admin": "Admin123456",
		},
	}))
	app.Get("/list", func(ctx *fiber.Ctx) error {
		return handlers.ListHandler(ctx, db)
	})
	app.Post("/edit/:Id", func(ctx *fiber.Ctx) error {
		return handlers.EditHandler(ctx, db)
	})
	logger.Fatal(app.Listen(":" + port))
}
