package domain

import "errors"

var (
	ErrURLNotFound         = errors.New("URL not found")
	ErrInternalServerError = errors.New("internal server error")
)
