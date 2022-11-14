package controllers

import (
	"yapp/core/internal/services/user"

	"github.com/gofiber/fiber/v2"
)

func PostLoginController(ctx *fiber.Ctx) error {

	return nil
}

func GetLoginController(ctx *fiber.Ctx) error {

	return nil
}

func PostRegisterController(ctx *fiber.Ctx) error {
	props := new(user.PostRegisterProps)
	ctx.BodyParser(props)

	result := user.RegisterUser(ctx.Context(), props)

	return ctx.Status(200).JSON(result)
}
