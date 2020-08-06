package requests

import "github.com/zoomi-raja/goweb/api/utils"

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (c Credentials) ValidateLoginRequest() utils.ErrorsArr {
	errArr := utils.ErrorArr{}
	if c.Email == "" {
		// errs["email"] = []string{"Email Field is missing"}
		errArr.AddError("email", "Email Field is missing")
	} else if !ValidEmail(c.Email) {
		errArr.AddError("email", "The email field should be a valid email address!")
	}
	if c.Password == "" {
		errArr.AddError("password", "Password is required!")
	} else if len(c.Password) < 6 || len(c.Password) > 40 {
		errArr.AddError("password", "The password field must be between 6-40 chars!")
	}
	return &errArr
}
