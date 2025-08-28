package auth

import (
	"fmt"
	"net/http"

	"github.com/FamousLuisin/agoraspace/internal/apperr"
	"github.com/FamousLuisin/agoraspace/internal/handler/user"
	"github.com/google/uuid"
)

func NewAuthService(repository user.UserRepository) AuthService {
	return &authService{
		repository: repository,
	}
}

type authService struct {
	repository user.UserRepository
}

type AuthService interface {
	SignUp(SignUpRequest) (string, *apperr.AppErr)
	SignIn(SignInRequest) (string, *apperr.AppErr)
}

func (s *authService) SignUp(ur SignUpRequest) (string, *apperr.AppErr) {
	if  err := PasswordValidation(ur.Password, ur.ConfirmPassword); err != nil{
		return "", apperr.NewAppError(fmt.Sprintf("password validation error: %s", err.Error()), apperr.ErrBadRequest, http.StatusBadRequest)
	}
	
	pw, err := PasswordEncoder(ur.Password)

	if err != nil {
		return "", apperr.NewAppError(fmt.Sprintf("password encryption error: %s", err.Error()), apperr.ErrInternalServer, http.StatusInternalServerError)
	}
	
	u := user.User{
		Id: uuid.New(),
		Name: ur.Name,
		Email: ur.Email,
		Password: pw,
		Username: ur.Username,
		DisplayName: ur.DisplayName,
		Birthday: ur.Birthday,
	}

	if err := s.repository.CreateUser(u); err != nil {
		return "", apperr.NewAppError(fmt.Sprintf("error inserting user into database: %s", err.Error()), apperr.ErrBadRequest, http.StatusBadRequest)
	}

	tokenString, err := GenerateToken(u.Id.String())
	
	if err != nil {
		return "", apperr.NewAppError(fmt.Sprintf("error generating token: %s", err.Error()), apperr.ErrInternalServer, http.StatusInternalServerError)
	}

	return tokenString, nil
}

func (s *authService) SignIn(ur SignInRequest) (string, *apperr.AppErr){
	
	us, err := s.repository.FindUserByEmail(ur.Email)
	
	if err != nil {
		return "", apperr.NewAppError(fmt.Sprintf("error geting user: %s", err.Error()), apperr.ErrBadRequest, http.StatusBadRequest)
	}

	if err := PassowrdMatch(us.Password, ur.Password); err != nil {
		return "", apperr.NewAppError(fmt.Sprintf("invalid password: %s", err.Error()), apperr.ErrBadRequest, http.StatusBadRequest)
	}
	
	if us.DeletedAt.Valid && !ur.Activate {
		return "", apperr.NewAppError("user deleted", apperr.ErrUnauthorized, http.StatusUnauthorized)
	}

	if !us.DeletedAt.Valid && ur.Activate {
		return "", apperr.NewAppError("user not deleted", apperr.ErrBadRequest, http.StatusBadRequest)
	}

	if us.DeletedAt.Valid && ur.Activate {
		if err := s.repository.ActivateUser(*us); err != nil {
			return "", apperr.NewAppError(fmt.Sprintf("error activating user: %s", err.Error()), apperr.ErrInternalServer, http.StatusInternalServerError)
		}
	}
	
	tokenString, err := GenerateToken(us.Id.String())
	
	if err != nil {
		return "", apperr.NewAppError(fmt.Sprintf("error generating token: %s", err.Error()), apperr.ErrInternalServer, http.StatusInternalServerError)
	}

	return tokenString, nil
}