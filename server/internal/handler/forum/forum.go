package forum

import (
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	Active   Status = "active"
	Deleted  Status = "deleted"
	Archived Status = "archived"
)

type Forum struct {
	Id          uuid.UUID `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	Status      Status    `db:"status"`
	IsPublic    bool      `db:"is_public"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
	Owner       uuid.UUID `db:"owner"`
}

type ForumRequest struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	IsPublic    bool      `json:"is_public"`
}

type ForumResponse struct {
	Id uuid.UUID `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	IsPublic bool `json:"is_public"`
	Status string `json:"status"`
	Owner uuid.UUID `json:"owner"`
}