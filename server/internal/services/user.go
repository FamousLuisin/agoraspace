package services

import (
	"fmt"
	"net/http"

	"github.com/FamousLuisin/agoraspace/internal/apperr"
	"github.com/FamousLuisin/agoraspace/internal/models"
	"github.com/FamousLuisin/agoraspace/internal/repository"
	"github.com/google/uuid"
)

func NewUserService(repository repository.UserRepository) UserService {
	return &userService{
		repository: repository,
	}
}

type userService struct {
	repository repository.UserRepository
}

type UserService interface {
	GetUserById(identifier string) (*models.UserDTO, *apperr.AppErr)
	GetUserByUsername(string) (*models.UserDTO, *apperr.AppErr)
	GetUsers(int, int) (*[]models.UserDTO, *apperr.AppErr)
	UpdateUser(models.UserDTO, string, string) *apperr.AppErr
	DeleteUser(string, string) *apperr.AppErr
}

func (s *userService) GetUserByUsername(username string) (*models.UserDTO, *apperr.AppErr){

	u, err := s.repository.FindUserByUsername(username)

	if err != nil {
		return nil, apperr.NewAppError(fmt.Sprintf("error getting user: %s", err.Error()), apperr.ErrBadRequest, http.StatusBadRequest)
	} 

	ur := models.UserDTO{
		Email: u.Email,
		Name: u.Name,
		Username: u.Username,
		Bio: u.Bio,
		Displayname: u.DisplayName,
	}

	return &ur, nil
}

func (s *userService) GetUsers(page, perPage int) (*[]models.UserDTO, *apperr.AppErr){
	users, err := s.repository.GetAllUsers(page, perPage)

	if err != nil {
		return nil, apperr.NewAppError(fmt.Sprintf("error getting users: %s", err.Error()), apperr.ErrBadRequest, http.StatusBadRequest)
	}

	var ur []models.UserDTO

	for _, user := range *users {
		u := models.UserDTO {
			Email: user.Email,
			Name: user.Name,
			Username: user.Username,
			Displayname: user.DisplayName,
			Bio: user.Bio,
		}

		ur = append(ur, u)
	}

	return &ur, nil
}

func (s userService) UpdateUser(userDTO models.UserDTO, identifier, username string) *apperr.AppErr {
	user, err := s.repository.FindUserByUsername(username)

	if err != nil {
		return apperr.NewAppError(fmt.Sprintf("error getting users: %s", err.Error()), apperr.ErrBadRequest, http.StatusBadRequest)
	}

	if user.Id.String() != identifier {
		return apperr.NewAppError("error validating identifier", apperr.ErrUnauthorized, http.StatusUnauthorized)
	}

	user.Bio = userDTO.Bio
	user.DisplayName = userDTO.Displayname
	user.Username = userDTO.Username
	user.Email = userDTO.Email
	user.Name = userDTO.Name
	
	if err := s.repository.UpdateUser(*user); err != nil{ 
		return apperr.NewAppError(fmt.Sprintf("error updating user: %s", err.Error()), apperr.ErrBadRequest, http.StatusBadRequest)
	}

	return nil
}

func (s *userService) DeleteUser(identifier, username string) *apperr.AppErr{
	user, err := s.repository.FindUserByUsername(username)

	if err != nil {
		return apperr.NewAppError(fmt.Sprintf("error getting users: %s", err.Error()), apperr.ErrBadRequest, http.StatusBadRequest)
	}

	if user.Id.String() != identifier {
		return apperr.NewAppError("error validating identifier", apperr.ErrUnauthorized, http.StatusUnauthorized)
	}

	if err := s.repository.DeleteUser(*user); err != nil {
		return apperr.NewAppError(fmt.Sprintf("error deleting user: %s", err.Error()), apperr.ErrInternalServer, http.StatusInternalServerError)
	}
	
	return nil
}

func (s *userService) GetUserById(identifier string) (*models.UserDTO, *apperr.AppErr) {
	_, err := uuid.Parse(identifier)

	if err != nil {
		return nil, apperr.NewAppError(fmt.Sprintf("uuid invalid: %s", err.Error()), apperr.ErrBadRequest, http.StatusBadRequest)
	}

	user, err := s.repository.FindUserById(identifier)

	if err != nil {
		return nil, apperr.NewAppError(fmt.Sprintf("error getting users: %s", err.Error()), apperr.ErrBadRequest, http.StatusBadRequest)
	}
	
	return &models.UserDTO{
		Email: user.Email,
		Name: user.Name,
		Username: user.Username,
		Displayname: user.DisplayName,
		Bio: user.Bio,
	}, nil
}