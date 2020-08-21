package requests

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (c Credentials) ValidateLoginRequest() ErrorsArr {
	varlidator := requestValidator{}
	varlidator.RequestFields = []field{
		field{
			required: true,
			name:     "email",
			isEmail:  true,
			value:    c.Email,
			message:  "Email Field is missing",
		},
		field{
			required: true,
			name:     "password",
			min:      6,
			max:      40,
			value:    c.Password,
			message:  "Password is required!",
		},
	}

	return varlidator.validate()
}
