package models

import (
	"net/mail"
	"strings"
	"unicode"
)

func ValidateEmail(email string) error {
	if email == "" {
		return ErrEmailRequired
	}

	email = strings.TrimSpace(email)
	_, err := mail.ParseAddress(email)
	if err != nil {
		return ErrEmailInvalid
	}

	if len(email) > 254 {
		return ErrEmailInvalid
	}

	return nil
}

func NormalizeEmail(email string) string {
	return strings.ToLower(strings.TrimSpace(email))
}

func ValidatePassword(password string) error {
	if password == "" {
		return ErrPasswordRequired
	}

	// Check minimum length
	if len(password) < 8 {
		return ErrPasswordTooShort
	}

	// Bcrypt has maximum length of 72 bytes
	if len(password) > 72 {
		return ErrPasswordTooLong
	}

	var (
		hasUpper   bool
		hasLower   bool
		hasNumber  bool
		hasSpecial bool
	)

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if !hasUpper || !hasLower || !hasNumber || !hasSpecial {
		return ErrPasswordTooWeak
	}

	return nil
}
