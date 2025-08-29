package services

import (
	"fmt"
	"net/http"

	"github.com/FamousLuisin/agoraspace/internal/apperr"
	"github.com/FamousLuisin/agoraspace/internal/models"
	"github.com/FamousLuisin/agoraspace/internal/repository"
	"github.com/google/uuid"
)

func NewMemberService(mr repository.MemberRepository, fr repository.ForumRepository) MemberService {
	return &memberService{
		memberRepository: mr,
		forumRepository: fr,
	}
}

type memberService struct {
	memberRepository repository.MemberRepository
	forumRepository  repository.ForumRepository
}

type MemberService interface {
	JoinForum(string, string) *apperr.AppErr
	LeaveForum(string, string) *apperr.AppErr
	FindForumsByMember(string) (*[]models.Member, *apperr.AppErr)
}

func (s *memberService) JoinForum(identifier, forumStr string) *apperr.AppErr{
	forumId, err := uuid.Parse(forumStr)

	if err != nil {
		return apperr.NewAppError(fmt.Sprintf("uuid invalid: %s", err.Error()), apperr.ErrBadRequest, http.StatusBadRequest)
	}
	
	f, err := s.forumRepository.FindForumById(forumStr)

	if err != nil {
		return apperr.NewAppError(fmt.Sprintf("forum not found: %s", err.Error()), apperr.ErrNotFound, http.StatusNotFound)
	}

	if !f.IsPublic || f.Status != models.Active {
		return apperr.NewAppError("unauthorized entry", apperr.ErrUnauthorized, http.StatusUnauthorized)
	}

	_, err = s.memberRepository.FindMember(identifier, forumStr)

	if err == nil {
		if err := s.memberRepository.ReturnForum(forumStr, identifier); err != nil {
			return apperr.NewAppError("error returning to the forum", apperr.ErrBadRequest, http.StatusBadRequest)
		}

		return nil
	}

	err = s.memberRepository.InsertMember(models.Member{
		UserId: uuid.MustParse(identifier),
		Role: models.Participant,
		ForumId: forumId,
	})

	if err != nil {
		return apperr.NewAppError(fmt.Sprintf("error adding member: %s", err.Error()), apperr.ErrBadRequest, http.StatusBadRequest)
	}

	return nil
}

func (s *memberService) LeaveForum(identifier, forumStr string) *apperr.AppErr {
	_, err := uuid.Parse(forumStr)

	if err != nil {
		return apperr.NewAppError(fmt.Sprintf("uuid invalid: %s", err.Error()), apperr.ErrBadRequest, http.StatusBadRequest)
	}

	_, err = s.memberRepository.FindMember(identifier, forumStr)

	if err != nil {
		return apperr.NewAppError(fmt.Sprintf("member not found: %s", err.Error()), apperr.ErrNotFound, http.StatusNotFound)
	}

	if err := s.memberRepository.LeaveForum(forumStr, identifier); err != nil {
		return apperr.NewAppError(fmt.Sprintf("error leaving the forum: %s", err.Error()), apperr.ErrBadRequest, http.StatusBadRequest)
	}

	return nil
}

func (s *memberService) FindForumsByMember(identifier string) (*[]models.Member, *apperr.AppErr){
	_, err := uuid.Parse(identifier)

	if err != nil {
		return nil, apperr.NewAppError(fmt.Sprintf("uuid invalid: %s", err.Error()), apperr.ErrBadRequest, http.StatusBadRequest)
	}

	members, err := s.memberRepository.FindForumsByMember(identifier)

	if err != nil {
		return nil, apperr.NewAppError(fmt.Sprintf("not found forums: %s", err.Error()), apperr.ErrNotFound, http.StatusNotFound)
	}

	return members, nil
}