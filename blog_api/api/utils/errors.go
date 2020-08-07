package utils

import (
	"fmt"
	"strings"
)

type ErrorsArr interface {
	AddError(string, string)
	GetErros() map[string]interface{}
	Error() string
	HasError() bool
}

var keyPairErrors map[string]interface{}

type ErrorArr struct {
	Errors map[string]interface{}
}

func (e *ErrorArr) AddError(key string, err string) {
	if e.Errors == nil {
		e.Errors = make(map[string]interface{})
	}
	e.Errors[key] = err
}

func (e *ErrorArr) Error() string {
	errors := []string{}
	for key, element := range e.Errors {
		errors = append(errors, fmt.Sprintf("key: %s => %s", key, element))
	}
	return strings.Join(errors, ",")
}
func (e *ErrorArr) GetErros() map[string]interface{} {
	return e.Errors
}

func (e *ErrorArr) HasError() bool {
	if len(e.Errors) > 0 {
		return true
	}
	return false
}
