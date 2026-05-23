package models

import (
	"errors"
)

var (
	// Generic errors
	ErrNotFound = errors.New("models: resource could not be found")

	// Email errors
	ErrEmailTaken       = errors.New("models: email address is already in use")
	ErrEmailRequired    = errors.New("models: email address is required")
	ErrEmailInvalid     = errors.New("models: email address is not valid")
	ErrEmailNotVerified = errors.New("models: email address not verified")

	// Password errors
	ErrPasswordRequired = errors.New("models: password is required")
	ErrPasswordTooShort = errors.New("models: password must be at least 8 characters")
	ErrPasswordTooLong  = errors.New("models: password must not exceed 72 characters")
	ErrPasswordTooWeak  = errors.New("models: password must contain at least one " +
		"uppercase letter, one lowercase letter, one number, and one special character")
)
