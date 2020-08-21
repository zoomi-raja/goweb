package requests

import (
	"fmt"
	"regexp"
)

var (
	regexpEmail *regexp.Regexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

type requestValidator struct {
	RequestFields []field
}
type field struct {
	required bool
	name     string
	min      int
	max      int
	message  string
	isEmail  bool
	value    string
}

func (r requestValidator) validate() *ErrorArr {
	errArr := ErrorArr{}
	if len(r.RequestFields) <= 0 {
		return &errArr
	}
	for _, field := range r.RequestFields {
		if field.required && len(field.value) <= 0 {
			errArr.AddError(field.name, field.message)
		} else if field.isEmail {
			if !ValidEmail(field.value) {
				errArr.AddError(field.name, "The "+field.name+" should be a valid email address!")
			}
		} else if field.min > 0 && len(field.value) < field.min {
			errArr.AddError(field.name, fmt.Sprintf("The %s length must be greater then %d chars!", field.name, field.min))
		} else if field.max > 0 && len(field.value) > field.max {
			errArr.AddError(field.name, fmt.Sprintf("The %s length must be less then %d chars!", field.name, field.max))
		}
	}
	return &errArr
}

func ValidEmail(email string) bool {
	result := regexpEmail.MatchString(email)
	return result
}
