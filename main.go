package main

import (
	"log"
	"time"

	"github.com/dhimweray222/test-be-rekadigital/config"
	"github.com/dhimweray222/test-be-rekadigital/controller"
	"github.com/dhimweray222/test-be-rekadigital/repository"
	"github.com/dhimweray222/test-be-rekadigital/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	time.Local = time.UTC
	serverConfig := config.NewServerConfig()
	db := config.NewPostgresPool()
	store := repository.NewStore(db)
	validate := validator.New()
	transactionQuery := repository.NewTransaction()
	transactionRepository := repository.NewTransactionRepository(store, transactionQuery)
	transactionMemberService := service.NewTransactionService(transactionRepository)
	transactionController := controller.NewTransactionController(validate, transactionMemberService)

	app := fiber.New(fiber.Config{BodyLimit: 10 * 1024 * 1024})
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "*",
		AllowHeaders:     "*",
		AllowCredentials: true,
	}))

	transactionController.Route(app)
	err := app.Listen(serverConfig.Host)
	if err != nil {
		log.Fatal(err)
	}
}
