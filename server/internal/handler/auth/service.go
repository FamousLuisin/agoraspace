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
	fmt.Println("Chegou no user service")

	if  err := PasswordValidation(ur.Password, ur.ConfirmPassword); err != nil{
		errMessage := fmt.Sprintf("password validation error: %s", err.Error())
		return "", apperr.NewAppError(errMessage, apperr.ErrBadRequest, http.StatusBadRequest)
	}
	
	pw, err := PasswordEncoder(ur.Password)

	if err != nil {
		errMessage := fmt.Sprintf("password encryption error: %s", err.Error())
		return "", apperr.NewAppError(errMessage, apperr.ErrInternalServer, http.StatusInternalServerError)
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
		errMessage := fmt.Sprintf("error inserting user into database: %s", err.Error())
		return "", apperr.NewAppError(errMessage, apperr.ErrBadRequest, http.StatusBadRequest)
	}

	tokenString, err := GenerateToken(u.Username)
	
	if err != nil {
		errMessage := fmt.Sprintf("error generating token: %s", err.Error())
		return "", apperr.NewAppError(errMessage, apperr.ErrInternalServer, http.StatusInternalServerError)
	}

	return tokenString, nil
}

func (s *authService) SignIn(ur SignInRequest) (string, *apperr.AppErr){
	
	us, err := s.repository.FindUserByEmail(ur.Email)
	
	if err != nil {
		errMessage := fmt.Sprintf("error geting user: %s", err.Error())
		return "", apperr.NewAppError(errMessage, apperr.ErrBadRequest, http.StatusBadRequest)
	}

	if err := PassowrdMatch(us.Password, ur.Password); err != nil {
		errMessage := fmt.Sprintf("invalid password: %s", err.Error())
		return "", apperr.NewAppError(errMessage, apperr.ErrBadRequest, http.StatusBadRequest)
	}
	
	tokenString, err := GenerateToken(us.Username)
	
	if err != nil {
		errMessage := fmt.Sprintf("error generating token: %s", err.Error())
		return "", apperr.NewAppError(errMessage, apperr.ErrInternalServer, http.StatusInternalServerError)
	}

	return tokenString, nil
}