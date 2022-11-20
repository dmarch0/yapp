package controllers

import (
	"yapp/core/internal/services/user"

	"github.com/gofiber/fiber/v2"
)

func PostLoginController(ctx *fiber.Ctx) error {
	props := new(user.LoginUserProps)
	err := ctx.BodyParser(props)
	if err != nil {
		return ctx.Status(401).Next()
	}

	result, token := user.LoginUser(ctx.Context(), props)
	if !result.Success {
		return ctx.Status(401).JSON(result)
	}
	ctx.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    token,
		Secure:   true,
		HTTPOnly: true,
	})

	return ctx.Status(200).JSON(result)
}

func GetLoginController(ctx *fiber.Ctx) error {

	return nil
}

func PostRegisterController(ctx *fiber.Ctx) error {
	props := new(user.RegisterUserProps)
	err := ctx.BodyParser(props)
	if err != nil {
		return ctx.Status(401).Next()
	}

	result := user.RegisterUser(ctx.Context(), props)
	if !result.Success {
		return ctx.Status(500).JSON(result)
	}

	return ctx.Status(200).JSON(result)
}
