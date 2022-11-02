package scheme_validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type IError struct {
	Field string
	Tag   string
	Value string
}

func ValidateBody[T any]() func(ctx *fiber.Ctx) error {
	Validator := validator.New()

	return func(ctx *fiber.Ctx) error {
		body := new(T)
		ctx.BodyParser(&body)

		err := Validator.Struct(body)
		if err != nil {
			errors := make([]*IError, 0)
			for _, err := range err.(validator.ValidationErrors) {
				var el IError
				el.Field = err.Field()
				el.Tag = err.Tag()
				el.Value = err.Param()
				errors = append(errors, &el)
			}
			return ctx.Status(fiber.StatusBadRequest).JSON(errors)
		}

		return ctx.Next()
	}
}
