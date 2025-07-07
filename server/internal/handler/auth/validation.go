package auth

import (
	"fmt"
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
		err := fmt.Errorf("password and confirmPassword cannot be different")
		return err
	}

	if len(password) < 8 {
		err := fmt.Errorf("passwords must contain 8 or more characters")
		return err
	}

	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasDigit := regexp.MustCompile(`\d`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[\W_]`).MatchString(password)

	if !(hasLower && hasUpper && hasDigit && hasSpecial) {
		err := fmt.Errorf("the password must contain at least one uppercase letter, one lowercase letter, one number, and one special character")
		return err
	}
	
	return nil
}

func PassowrdMatch(hash, password string) error{
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}