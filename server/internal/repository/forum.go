package repository

import (
	"time"

	"github.com/FamousLuisin/agoraspace/internal/db"
	"github.com/FamousLuisin/agoraspace/internal/models"
	"github.com/jmoiron/sqlx"
)

const (
	insertForumQuery = `
		WITH inserted_forum AS (
			INSERT INTO tb_forums (id, title, description, is_public, owner)
			VALUES ($1, $2, $3, $4, $5)
			RETURNING id, owner
		)
		INSERT INTO tb_members (user_id, forum_id, role)
		SELECT owner, id, $6
		FROM inserted_forum
	`
	getAllForumsQuery = `
		SELECT * FROM tb_forums
		OFFSET $1 LIMIT $2
	`
	findForumByIdQuery = `
		SELECT * FROM tb_forums WHERE id = $1
	`
	findMemberByForum = `
		SELECT m.*, u.username, f.title FROM tb_members m
		JOIN tb_forums f ON m.forum_id = f.id
		JOIN tb_users u ON u.id = m.user_id
		WHERE m.forum_id = $1
	`
	updateForumQuery = `
		UPDATE tb_forums
		SET title = $1, description = $2, is_public = $3, updated_at = $4
		WHERE id = $5 AND owner = $6
	`
	deleteForumQuery = `
		UPDATE tb_forums
		SET status = $1
		WHERE id = $2 AND owner = $3
	`
)

func NewForumRepository(db *db.Database) ForumRepository {
	return &forumRepository{
		db: db.Db,
	}
}

type forumRepository struct{
	db *sqlx.DB
}

type ForumRepository interface{
	CreateForum(models.Forum) error
	GetAllForums(int, int) (*[]models.Forum, error)
	FindForumById(string) (*models.Forum, error)
	UpdateForum(models.Forum) error
	DeleteForum(models.Forum) error
}

func (r *forumRepository) CreateForum(f models.Forum) error {
	
	if _, err := r.db.Exec(insertForumQuery, f.Id, f.Title, f.Description, f.IsPublic, f.Owner, "owner"); err != nil {
		return err
	}

	return nil
}

func (r *forumRepository) GetAllForums(page, perPage int) (*[]models.Forum, error){
	var forums []models.Forum 

	if err := r.db.Select(&forums, getAllForumsQuery, page, perPage); err != nil {
		return nil, err
	}
	
	return &forums, nil
}

func (r *forumRepository) FindForumById(forumId string) (*models.Forum, error) {
	var f models.Forum
	
	if err := r.db.Get(&f, findForumByIdQuery, forumId); err != nil {
		return nil, err
	}

	return  &f, nil
}

func (r *forumRepository) UpdateForum(forum models.Forum) error {
	_, err := r.db.Exec(updateForumQuery, forum.Title, forum.Description, forum.IsPublic, time.Now(), forum.Id, forum.Owner)

	if err != nil {
		return err
	}

	return nil
}

func (r *forumRepository) DeleteForum(forum models.Forum) error {
	_, err := r.db.Exec(deleteForumQuery, models.Deleted, forum.Id, forum.Owner)

	if err != nil {
		return err
	}

	return nil
}

func (r *forumRepository) FindMembersByForum(forumId string) (*[]models.Member, error){
	var members []models.Member

	if err := r.db.Select(&members, findMemberByForum, forumId); err != nil{
		return nil, err
	}

	return &members, nil
}