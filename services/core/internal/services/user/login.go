package user

import (
	"context"
	"yapp/core/internal/crypt"
	"yapp/core/internal/db"
	"yapp/core/internal/http/scheme_validation"

	"golang.org/x/crypto/bcrypt"
)

type LoginUserProps scheme_validation.PostLoginBody

type UserData struct {
	Id string `json:"id"`
}

func LoginUser(ctx context.Context, props *LoginUserProps) (string, UserData, error) {
	connection := db.GetDbConnection()

	var password_hash []byte
	var id string
	connection.QueryRow(ctx, `SELECT password, id FROM users WHERE email=$1`, props.Email).Scan(&password_hash, &id)
	if err := bcrypt.CompareHashAndPassword(password_hash, []byte(props.Password)); err != nil {
		return "", UserData{}, err
	}

	token := crypt.GenerateToken(16)
	if _, err := connection.Exec(ctx, `DELETE FROM user_tokens WHERE user_id=$1`, id); err != nil {
		return "", UserData{}, err
	}

	_, tokenError := connection.Exec(ctx, `
		INSERT INTO user_tokens(user_id, token)
		VALUES(
			$1,
			$2
		);
	`, id, token)
	if tokenError != nil {
		return "", UserData{}, tokenError
	}

	return token, UserData{id}, nil
}
