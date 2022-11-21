package user

import (
	"context"
	"log"
	"yapp/core/internal/crypt"
	"yapp/core/internal/db"
	"yapp/core/internal/http/scheme_validation"
)

type RegisterUserProps scheme_validation.PostRegisterBody

func RegisterUser(ctx context.Context, props *RegisterUserProps) error {
	connection := db.GetDbConnection()
	hash, hash_err := crypt.HashPassword(props.Password)
	if hash_err != nil {
		return hash_err
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
		return err
	}

	return nil
}
