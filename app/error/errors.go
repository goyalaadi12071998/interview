package errorclass

import "strings"

type Error struct {
	name        string
	code        string
	description string
}

type IError interface {
	StatusCode() int
	Name() string
	Code() string
	Description() string
	Wrap() string
}

func (e *Error) StatusCode() int {
	if strings.Contains(e.code, "INTERNAL_SERVER") {
		return 500
	}
	return 400
}

func (e *Error) Name() string {
	return e.name
}

func (e *Error) Code() string {
	return e.code
}

func (e *Error) Description() string {
	return e.description
}

func (e *Error) Wrap(des string) *Error {
	e.description = des
	return e
}

func NewError(errorcode string) *Error {
	return errorList[errorcode]
}
