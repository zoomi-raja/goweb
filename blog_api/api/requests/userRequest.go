package requests

type User struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
}

func (u User) ValidateUserCreate() ErrorsArr {
	varlidator := requestValidator{}
	varlidator.RequestFields = []field{
		field{
			required: true,
			name:     "email",
			isEmail:  true,
			value:    u.Email,
			message:  "Email is missing",
		},
		field{
			required: true,
			name:     "username",
			min:      2,
			max:      40,
			value:    u.Username,
			message:  "Username is required!",
		},
		field{
			required: true,
			name:     "password",
			min:      6,
			max:      40,
			value:    u.Password,
			message:  "Password is required!",
		},
	}

	return varlidator.validate()
}
