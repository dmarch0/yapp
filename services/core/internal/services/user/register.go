package user

import (
	"context"
	"log"
	"yapp/core/internal/crypt"
	"yapp/core/internal/db"
	"yapp/core/internal/http/scheme_validation"
)

type RegisterUserResult struct {
	Success bool `json:"success"`
}

type PostRegisterProps scheme_validation.PostRegisterBody

func RegisterUser(ctx context.Context, props *PostRegisterProps) RegisterUserResult {
	connection := db.GetDbConnection()
	_, err := connection.Exec(ctx, `
		INSERT INTO users(email, password)
		VALUES(
			$1,
			$2
		);
	`, props.Email, crypt.HashPassword(props.Password))
	if err != nil {
		log.Println("Error inserting new user", err)
		return RegisterUserResult{
			Success: false,
		}
	}

	return RegisterUserResult{
		Success: true,
	}
}
