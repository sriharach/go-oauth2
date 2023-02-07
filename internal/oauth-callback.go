package internal

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type loginCallbackBody struct {
	Code  string `json:"code"`
	State string `json:"state"`
}

func LoginCallback(c *fiber.Ctx) error {
	request := new(loginCallbackBody)

	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	config := Oauth2()
	token, err := config.Exchange(c.Context(), request.Code)
	fmt.Println(token)

	if err != nil {
		fmt.Println("CannotGetGoogleToken: %s", err.Error())
		return c.Status(fiber.StatusServiceUnavailable).JSON(err.Error())
	}

	// id := token.Extra("id_token").(string)
	return c.JSON(token)
}
