package controllers

import (
	"yapp/core/internal/services/user"

	"github.com/gofiber/fiber/v2"
)

type PostLoginResult struct {
	Success bool          `json:"success"`
	Data    user.UserData `json:"data"`
	Error   string        `json:"error"`
}

func PostLoginController(ctx *fiber.Ctx) error {
	props := new(user.LoginUserProps)
	err := ctx.BodyParser(props)
	if err != nil {
		return ctx.Status(401).Next()
	}

	token, user_data, err := user.LoginUser(ctx.Context(), props)
	if err != nil {
		return ctx.Status(401).JSON(PostLoginResult{
			Success: false,
			Error:   err.Error(),
		})
	}
	ctx.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    token,
		Secure:   true,
		HTTPOnly: true,
	})

	return ctx.Status(200).JSON(PostLoginResult{
		Success: true,
		Error:   err.Error(),
		Data:    user_data,
	})
}

func GetLoginController(ctx *fiber.Ctx) error {

	return nil
}

type PostRegisterResult struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

func PostRegisterController(ctx *fiber.Ctx) error {
	props := new(user.RegisterUserProps)
	err := ctx.BodyParser(props)
	if err != nil {
		return ctx.Status(401).Next()
	}

	registerError := user.RegisterUser(ctx.Context(), props)
	if registerError != nil {
		return ctx.Status(500).JSON(PostRegisterResult{
			Success: false,
			Error:   registerError.Error(),
		})
	}

	return ctx.Status(200).JSON(PostRegisterResult{
		Success: true,
	})
}
