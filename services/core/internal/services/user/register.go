package user

import (
	"context"
	"log"
	"yapp/core/internal/crypt"
	"yapp/core/internal/db"
	"yapp/core/internal/http/scheme_validation"
)

type RegisterUserResult struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

type RegisterUserProps scheme_validation.PostRegisterBody

func RegisterUser(ctx context.Context, props *RegisterUserProps) RegisterUserResult {
	connection := db.GetDbConnection()
	hash, hash_err := crypt.HashPassword(props.Password)
	if hash_err != nil {
		return RegisterUserResult{
			Success: false,
			Error:   hash_err.Error(),
		}
	}

	_, err := connection.Exec(ctx, `
		INSERT INTO users(email, password)
		VALUES(
			$1,
			$2
		);
	`, props.Email, hash)
	if err != nil {
		log.Println("Error inserting new user", err)
		return RegisterUserResult{
			Success: false,
			Error:   err.Error(),
		}
	}

	return RegisterUserResult{
		Success: true,
	}
}
