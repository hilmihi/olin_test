package controller

import (
	"olin_test/adapter/request"
	"olin_test/usecase"

	"github.com/gofiber/fiber/v2"
)

type ITransactionHandler interface {
	Soal1(c *fiber.Ctx) error
	Soal2(c *fiber.Ctx) error
	Soal3(c *fiber.Ctx) error
}

type transactionHandler struct {
	usecase usecase.GoAssestmentUseCaseInterface
}

func NewAssestmentHandler(u usecase.GoAssestmentUseCaseInterface) ITransactionHandler {
	return &transactionHandler{
		usecase: u,
	}
}

func (t *transactionHandler) Soal1(c *fiber.Ctx) error {
	var soal1Req request.Soal1Request

	err := c.BodyParser(&soal1Req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	res, err1 := t.usecase.Soal1(soal1Req.Nums, soal1Req.Target)
	if err1 != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err1)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (t *transactionHandler) Soal2(c *fiber.Ctx) error {
	var soal2Req request.Soal2Request

	err := c.BodyParser(&soal2Req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	res, err1 := t.usecase.Soal2(soal2Req.Nums)
	if err1 != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err1)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (t *transactionHandler) Soal3(c *fiber.Ctx) error {
	var soal3Req request.Soal3Request

	err := c.BodyParser(&soal3Req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	res, err1 := t.usecase.Soal3(soal3Req.Str, soal3Req.Words)
	if err1 != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err1)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
