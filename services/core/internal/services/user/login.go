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

type LoginUserResult struct {
	Success bool     `json:"success"`
	Data    UserData `json:"data"`
	Error   string   `json:"error"`
}

func LoginUser(ctx context.Context, props *LoginUserProps) (LoginUserResult, string) {
	connection := db.GetDbConnection()

	var password_hash []byte
	var id string
	connection.QueryRow(ctx, `SELECT password, id FROM users WHERE email=$1`, props.Email).Scan(&password_hash, &id)
	err := bcrypt.CompareHashAndPassword(password_hash, []byte(props.Password))
	if err != nil {
		return LoginUserResult{
			Success: false,
			Error:   err.Error(),
		}, ""
	}
	token := crypt.GenerateToken(16)
	_, deleteError := connection.Exec(ctx, `DELETE FROM user_tokens WHERE user_id=$1`, id)
	if deleteError != nil {
		return LoginUserResult{
			Success: false,
			Error:   deleteError.Error(),
		}, ""
	}
	_, tokenError := connection.Exec(ctx, `
		INSERT INTO user_tokens(user_id, token)
		VALUES(
			$1,
			$2
		);
	`, id, token)
	if tokenError != nil {
		return LoginUserResult{
			Success: false,
			Error:   tokenError.Error(),
		}, ""
	}

	return LoginUserResult{
		Success: true,
		Data: UserData{
			Id: id,
		},
	}, token
}
