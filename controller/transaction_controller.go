package controller

import (
	"log"

	"github.com/dhimweray222/test-be-rekadigital/config"
	"github.com/dhimweray222/test-be-rekadigital/exception"
	"github.com/dhimweray222/test-be-rekadigital/model/web"
	"github.com/dhimweray222/test-be-rekadigital/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type TransactionController interface {
	Route(app *fiber.App)
	CreateTransaction(ctx *fiber.Ctx) error
}

type transactionController struct {
	validate           *validator.Validate
	transactionService service.TransactionService
}

func NewTransactionController(validate *validator.Validate, transactionService service.TransactionService) TransactionController {
	return &transactionController{

		validate:           validate,
		transactionService: transactionService,
	}
}
func (controller *transactionController) Route(app *fiber.App) {
	api := app.Group(config.EndpointPrefix)

	api.Post("/",
		controller.CreateTransaction,
	)

	api.Get("/",
		controller.FindAll,
	)
}
func (controller *transactionController) CreateTransaction(ctx *fiber.Ctx) error {

	var request web.CreateTransactionRequest
	err := ctx.BodyParser(&request)
	if err != nil {
		log.Println(err)
		return exception.ErrorHandler(ctx, err)
	}

	err = controller.validate.Struct(request)
	if err != nil {
		return exception.ErrValidateBadRequest(err.Error(), request)
	}
	transactionResponse, err := controller.transactionService.CreateTransaction(ctx.Context(), request)
	if err != nil {
		return exception.ErrorHandler(ctx, err)
	}

	return ctx.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:    fiber.StatusOK,
		Status:  true,
		Message: "success",
		Data:    transactionResponse,
	})
}

func (controller *transactionController) FindAll(ctx *fiber.Ctx) error {

	menu := ctx.Query("menu")
	customerId := ctx.Query("customer_id")
	transactionResponses, totalData, err := controller.transactionService.FindAll(ctx.Context(), menu, customerId)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(web.WebResponseCount{
		Code:      fiber.StatusOK,
		Status:    true,
		TotalData: totalData,
		Message:   "success",
		Data:      transactionResponses,
	})
}
