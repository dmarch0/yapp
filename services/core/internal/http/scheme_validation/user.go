package scheme_validation

type PostRegisterBody struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
