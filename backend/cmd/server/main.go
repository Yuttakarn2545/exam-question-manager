package main

import (
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"

	"exam-question-manager/internal/handler"
	"exam-question-manager/internal/repository"
	"exam-question-manager/internal/service"
)

func main() {
	// .env is optional — local dev convenience only, ignore if it's missing.
	_ = godotenv.Load()

	dataFile := os.Getenv("DATA_FILE")
	if dataFile == "" {
		dataFile = "data/exams.json"
	}

	repo, err := repository.NewJSONFileExamRepository(dataFile)
	if err != nil {
		log.Fatalf("failed to load data file %q: %v", dataFile, err)
	}

	examService := service.NewExamService(repo)
	examHandler := handler.NewExamHandler(examService)

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOriginsFunc: func(origin string) bool {
			return strings.HasPrefix(origin, "http://localhost:")
		},
		AllowMethods: "GET,POST,DELETE",
	}))

	v1 := app.Group("/api/v1")
	examHandler.Register(v1)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(app.Listen(":" + port))
}
