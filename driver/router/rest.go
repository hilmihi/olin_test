package router

import (
	"olin_test/adapter/controller"
	"olin_test/config"
	"olin_test/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Init(useCase usecase.AppUseCases) {
	router := fiber.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:9000",
		AllowHeaders:     "*",
		AllowMethods:     "GET, POST, PATCH, OPTIONS, PUT, DELETE",
		ExposeHeaders:    "Content-Length",
		AllowCredentials: true,
	}),
	)

	goassestment := controller.NewAssestmentHandler(useCase.TransactionUseCase)

	router.Post("/soal-1", goassestment.Soal1)
	router.Post("/soal-2", goassestment.Soal2)
	router.Post("/soal-3", goassestment.Soal3)

	port := config.GetEnv("PORT")
	if port == "" {
		port = "9000"
	}

	err := router.Listen(":" + port)
	if err != nil {

	}
}
