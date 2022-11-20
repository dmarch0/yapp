package routes

import (
	"yapp/core/internal/http/controllers"
	"yapp/core/internal/http/scheme_validation"

	"github.com/gofiber/fiber/v2"
)

func AddUserEndpoints(router fiber.Router) {
	router.Post(
		"/login",
		scheme_validation.ValidateBody[scheme_validation.PostLoginBody](),
		controllers.PostLoginController,
	)

	router.Get("/login", controllers.GetLoginController)

	router.Post(
		"/register",
		scheme_validation.ValidateBody[scheme_validation.PostRegisterBody](),
		controllers.PostRegisterController,
	)

}
