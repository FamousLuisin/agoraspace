package auth

import (
	"time"
)

type SignUpRequest struct {
	Name               string    `json:"name"`
	Email              string    `json:"email"`
	Password           string    `json:"password"`
	ConfirmPassword    string    `json:"confirmPassword"`
	Username           string    `json:"username"`
	DisplayName        string    `json:"displayName"`
	BirthdayString     string    `json:"birthday"`
	Birthday           time.Time `json:"-"`
}

type AuthResponse struct {
	Token string `json:"token"`
}

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Activate bool   `json:"activate"`
}