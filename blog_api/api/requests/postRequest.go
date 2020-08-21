package requests

/*Post struct type to hold request*/
type Post struct {
	Title string `json:"title"`
	Intro string `json:"intro"`
	Body  string `json:"body"`
}

/*ValidatePostCreate function to validate post before even to try saving post*/
func (p Post) ValidatePostCreate() ErrorsArr {
	varlidator := requestValidator{}
	varlidator.RequestFields = []field{
		field{
			required: true,
			name:     "title",
			value:    p.Title,
			message:  "Title is missing",
		},
		field{
			required: true,
			name:     "intro",
			min:      10,
			max:      300,
			value:    p.Intro,
			message:  "Intro is required!",
		},
		field{
			required: true,
			name:     "body",
			min:      300,
			value:    p.Body,
			message:  "Body is required!",
		},
	}

	return varlidator.validate()
}
