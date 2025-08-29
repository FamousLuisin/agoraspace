package services

import (
	"fmt"
	"net/http"

	"github.com/FamousLuisin/agoraspace/internal/apperr"
	"github.com/FamousLuisin/agoraspace/internal/models"
	"github.com/FamousLuisin/agoraspace/internal/repository"
	"github.com/google/uuid"
)

func NewForumService(repository repository.ForumRepository) ForumService {
	return &forumService{
		repository: repository,
	}
}

type forumService struct {
	repository repository.ForumRepository
}

type ForumService interface {
	CreateForum(models.ForumRequest, string) *apperr.AppErr
	GetAllForums(int, int) (*[]models.ForumResponse, *apperr.AppErr)
	GetForumById(string) (*models.ForumResponse, *apperr.AppErr)
	UpdateForum(models.ForumRequest, string, string) (*models.ForumResponse, *apperr.AppErr)
	DeleteForum(string, string) *apperr.AppErr
}

func (s *forumService) CreateForum(forumR models.ForumRequest, identifier string) *apperr.AppErr{
	forum := models.Forum{
		Id: uuid.New(),
		Title: forumR.Title,
		Description: forumR.Description,
		IsPublic: forumR.IsPublic,
		Owner: uuid.MustParse(identifier),
	}

	if err := s.repository.CreateForum(forum); err != nil {
		return apperr.NewAppError(fmt.Sprintf("error creating forum: %s", err.Error()), apperr.ErrBadRequest, http.StatusBadRequest)
	}

	return nil
}

func (s *forumService) GetAllForums(page, perPage int) (*[]models.ForumResponse, *apperr.AppErr){
	forums, err := s.repository.GetAllForums(page, perPage)

	if err != nil {
		return nil, apperr.NewAppError(fmt.Sprintf("error getting forums: %s", err.Error()), apperr.ErrNotFound, http.StatusNotFound)
	}

	var forumsResponse []models.ForumResponse

	for _, f := range *forums {
		forumR := models.ForumResponse{
			Id: f.Id,
			Title: f.Title,
			Description: f.Description,
			IsPublic: f.IsPublic,
			Status: string(f.Status),
			Owner: f.Owner,
		}

		forumsResponse = append(forumsResponse, forumR)
	}

	return &forumsResponse, nil
}

func (s *forumService) GetForumById(forumId string) (*models.ForumResponse, *apperr.AppErr) {
	_, err := uuid.Parse(forumId)

	if err != nil {
		return nil, apperr.NewAppError(fmt.Sprintf("uuid invalid: %s", err.Error()), apperr.ErrBadRequest, http.StatusBadRequest)
	}

	forum, err := s.repository.FindForumById(forumId)

	if err != nil {
		return nil, apperr.NewAppError(fmt.Sprintf("forum not found: %s", err.Error()), apperr.ErrNotFound, http.StatusNotFound)
	}

	return &models.ForumResponse{
		Id: forum.Id,
		Title: forum.Title,
		Description: forum.Description,
		IsPublic: forum.IsPublic,
		Status: string(forum.Status),
		Owner: forum.Owner,
	}, nil
}

func (s *forumService) UpdateForum(forumR models.ForumRequest, identifier, forumId string) (*models.ForumResponse, *apperr.AppErr) {

	_, err := uuid.Parse(forumId)

	if err != nil {
		return nil, apperr.NewAppError(fmt.Sprintf("uuid invalid: %s", err.Error()), apperr.ErrBadRequest, http.StatusBadRequest)
	}

	forum, err := s.repository.FindForumById(forumId)

	if err != nil {
		return nil, apperr.NewAppError(fmt.Sprintf("forum not found: %s", err.Error()), apperr.ErrNotFound, http.StatusNotFound)
	}

	if forum.Owner.String() != identifier {
		return nil, apperr.NewAppError("invalid operation", apperr.ErrUnauthorized, http.StatusUnauthorized)
	}

	forum.Title = forumR.Title
	forum.Description = forumR.Description
	forum.IsPublic = forumR.IsPublic

	err = s.repository.UpdateForum(*forum)

	if err != nil {
		return nil, apperr.NewAppError(fmt.Sprintf("error updating forum: %s", err.Error()), apperr.ErrBadRequest, http.StatusBadRequest)
	}

	return &models.ForumResponse{
		Id: forum.Id,
		Title: forum.Title,
		Description: forum.Description,
		IsPublic: forum.IsPublic,
		Status: string(forum.Status),
		Owner: forum.Owner,
	}, nil
}

func (s *forumService) DeleteForum(identifier, forumId string) *apperr.AppErr {

	_, err := uuid.Parse(forumId)

	if err != nil {
		return apperr.NewAppError(fmt.Sprintf("uuid invalid: %s", err.Error()), apperr.ErrBadRequest, http.StatusBadRequest)
	}

	forum, err := s.repository.FindForumById(forumId)

	if err != nil {
		return apperr.NewAppError(fmt.Sprintf("forum not found: %s", err.Error()), apperr.ErrNotFound, http.StatusNotFound)
	}

	if forum.Owner.String() != identifier {
		return apperr.NewAppError("invalid operation", apperr.ErrUnauthorized, http.StatusUnauthorized)
	}

	err = s.repository.DeleteForum(*forum)

	if err != nil {
		return apperr.NewAppError(fmt.Sprintf("error updating forum: %s", err.Error()), apperr.ErrBadRequest, http.StatusBadRequest)
	}

	return nil
}