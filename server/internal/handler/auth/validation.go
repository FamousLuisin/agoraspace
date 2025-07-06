package auth

import (
	"errors"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func PasswordEncoder(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func PasswordValidation(password, confirmPassword string) error {
	
	if password != confirmPassword{
		err := errors.New("password and confirmPassword cannot be different")
		return err
	}

	if len(password) < 8 {
		err := errors.New("passwords must contain 8 or more characters")
		return err
	}

	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasDigit := regexp.MustCompile(`\d`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[\W_]`).MatchString(password)

	if !(hasLower && hasUpper && hasDigit && hasSpecial) {
		err := errors.New("the password must contain at least one uppercase letter, one lowercase letter, one number, and one special character")
		return err
	}
	
	return nil
}

func PassowrdMatch(hash, password string) error{
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}