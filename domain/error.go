package domain

import "errors"

var (
	ErrInternalServerError = errors.New("Internal Server Error")
	ErrNotFound            = errors.New("Not Found")
	ErrForbidden           = errors.New("Forbidden")
	ErrUnauthorized        = errors.New("Unauthorized")
	ErrBadRequest          = errors.New("Bad Request")
	StatusCreated          = errors.New("Created")
	ErrUnknownType         = errors.New("Unknown Type")
	ErrAlreadyExistsPlan   = errors.New("Already Exists")
)
