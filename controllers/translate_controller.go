package controllers

import (
	"context"
	"fmt"
	"github.com/danieljvx/meli-morse/commons"
	"github.com/danieljvx/meli-morse/requests"
	"github.com/danieljvx/meli-morse/responses"
	"net/http"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func DecodeBits(c *fiber.Ctx) error {
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var translate requests.TranslateRequest
	defer cancel()

	if err := c.BodyParser(&translate); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.TranslateResponse{Status: http.StatusBadRequest, Message: err.Error()})
	}

	if validationErr := validate.Struct(&translate); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.TranslateResponse{Status: http.StatusBadRequest, Message: validationErr.Error()})
	}

	t := commons.NewTranslate()

	morseCode, err := t.DecodeBits2Morse(strings.NewReader(translate.Text))
	if err != nil {
		fmt.Println(err)
	}

	return c.Status(http.StatusCreated).JSON(responses.TranslateResponse{Status: http.StatusCreated, Message: string(morseCode) })
}

func ToText(c *fiber.Ctx) error {
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var translate requests.TranslateRequest
	defer cancel()

	if err := c.BodyParser(&translate); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.TranslateResponse{Status: http.StatusBadRequest, Message: err.Error()})
	}

	if validationErr := validate.Struct(&translate); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.TranslateResponse{Status: http.StatusBadRequest, Message: validationErr.Error()})
	}

	t := commons.NewTranslate()
	textCode, err := t.Translate2Human(strings.NewReader(translate.Text))
	if err != nil {
		fmt.Println(err)
	}

	return c.Status(http.StatusCreated).JSON(responses.TranslateResponse{Status: http.StatusCreated, Message: string(textCode) })
}

func ToMorse(c *fiber.Ctx) error {
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var translate requests.TranslateRequest
	defer cancel()

	if err := c.BodyParser(&translate); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.TranslateResponse{Status: http.StatusBadRequest, Message: err.Error()})
	}

	if validationErr := validate.Struct(&translate); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.TranslateResponse{Status: http.StatusBadRequest, Message: validationErr.Error()})
	}

	t := commons.NewTranslate()
	morseCode, err := t.Translate2Morse(strings.NewReader(translate.Text))
	if err != nil {
		fmt.Println(err)
	}

	return c.Status(http.StatusCreated).JSON(responses.TranslateResponse{Status: http.StatusCreated, Message: string(morseCode) })
}