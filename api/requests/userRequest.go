package requests

import (
	"github.com/zoomi-raja/goweb/api/utils"
)

type User struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
}

func (u User) ValidateUserCreate() utils.ErrorsArr {
	errArr := utils.ErrorArr{}
	if u.Email == "" {
		// errs["email"] = []string{"Email Field is missing"}
		errArr.AddError("email", "Email Field is missing")
	} else if !ValidEmail(u.Email) {
		errArr.AddError("email", "The email field should be a valid email address!")
	}

	// check if the name empty
	if u.Username == "" {
		errArr.AddError("name", "The name is required!")
	}
	// check the name field is between 3 to 120 chars
	if len(u.Username) < 2 || len(u.Username) > 40 {
		errArr.AddError("name", "The name field must be between 2-40 chars!")
	}
	// check if the password empty
	if u.Password == "" {
		errArr.AddError("password", "Password is required!")
	} else if len(u.Password) < 6 || len(u.Password) > 40 {
		errArr.AddError("password", "The password field must be between 6-40 chars!")
	}

	return &errArr
}
