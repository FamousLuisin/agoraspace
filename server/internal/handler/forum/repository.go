package forum

import (
	"time"

	"github.com/FamousLuisin/agoraspace/internal/db"
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
	CreateForum(Forum) error
	GetAllForums(int, int) (*[]Forum, error)
	FindForumById(string) (*Forum, error)
	UpdateForum(Forum) error
	DeleteForum(Forum) error
}

func (r *forumRepository) CreateForum(f Forum) error {
	
	if _, err := r.db.Exec(insertForumQuery, f.Id, f.Title, f.Description, f.IsPublic, f.Owner, "owner"); err != nil {
		return err
	}

	return nil
}

func (r *forumRepository) GetAllForums(page, perPage int) (*[]Forum, error){
	var forums []Forum 

	if err := r.db.Select(&forums, getAllForumsQuery, page, perPage); err != nil {
		return nil, err
	}
	
	return &forums, nil
}

func (r *forumRepository) FindForumById(forumId string) (*Forum, error) {
	var f Forum
	
	if err := r.db.Get(&f, findForumByIdQuery, forumId); err != nil {
		return nil, err
	}

	return  &f, nil
}

func (r *forumRepository) UpdateForum(forum Forum) error {
	_, err := r.db.Exec(updateForumQuery, forum.Title, forum.Description, forum.IsPublic, time.Now(), forum.Id, forum.Owner)

	if err != nil {
		return err
	}

	return nil
}

func (r *forumRepository) DeleteForum(forum Forum) error {
	_, err := r.db.Exec(deleteForumQuery, Deleted, forum.Id, forum.Owner)

	if err != nil {
		return err
	}

	return nil
}