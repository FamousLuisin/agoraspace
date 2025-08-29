package repository

import (
	"github.com/FamousLuisin/agoraspace/internal/db"
	"github.com/FamousLuisin/agoraspace/internal/models"
	"github.com/jmoiron/sqlx"
)

const (
	insertMemberQuery = `
		INSERT INTO tb_members
		(user_id, forum_id, role)
		VALUES ($1, $2, $3)
	`
	activeForumQuery = `
		UPDATE tb_members
		SET active = $1
		WHERE user_id = $2 AND forum_id = $3	
	`
	findMemberQuery = `
		SELECT * FROM tb_members
		WHERE user_id = $1 AND forum_id = $2
	`
	findForumsByMemberQuery = `
		SELECT m.*, u.username, f.title FROM tb_members m 
		JOIN tb_forums f ON f.id = m.forum_id 
		JOIN tb_users u ON u.id = m.user_id 
		WHERE m.user_id = $1
	`
)

func NewMemberRepository(db *db.Database) MemberRepository {
	return &memberRepository{
		db: db.Db,
	}
}

type memberRepository struct {
	db *sqlx.DB
}

type MemberRepository interface {
	InsertMember(models.Member) error
	LeaveForum(string, string) error
	ReturnForum(string, string) error
	FindMember(string, string) (*models.Member, error)
	FindForumsByMember(string) (*[]models.Member, error)
}

func (r *memberRepository) InsertMember(member models.Member) error {
	_, err := r.db.Exec(insertMemberQuery, member.UserId, member.ForumId, member.Role)

	if err != nil {
		return err
	}

	return nil
}

func (r *memberRepository) LeaveForum(forumId, userId string) error{
	_, err := r.db.Exec(activeForumQuery, false, userId, forumId)

	if err != nil {
		return err
	}

	return nil
}

func (r *memberRepository) ReturnForum(forumId, userId string) error{
	_, err := r.db.Exec(activeForumQuery, true, userId, forumId)

	if err != nil {
		return err
	}

	return nil
}

func (r *memberRepository) FindMember(userId, forumId string) (*models.Member, error) {
	var m models.Member

	if err := r.db.Get(&m, findMemberQuery, userId, forumId); err != nil {
		return nil, err
	}

	return &m, nil
}

func (r *memberRepository) FindForumsByMember(userId string) (*[]models.Member, error) {
	var m []models.Member

	if err := r.db.Select(&m, findForumsByMemberQuery, userId); err != nil {
		return nil, err
	}

	return &m, nil
}