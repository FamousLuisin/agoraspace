package user

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Role string

const (
	Admin  Role = "admin" 
	Common Role = "common"
)

type User struct {
	Id           uuid.UUID    `db:"id"`
	Email        string       `db:"email"`
	Name         string       `db:"name"`
	Username     string       `db:"username"`
	DisplayName  string       `db:"display_name"`
	Bio          string	      `db:"bio"`
	Birthday     time.Time    `db:"birthday"`
	Password     string       `db:"password"`
	Role         Role         `db:"role"`
	CreatedAt    time.Time    `db:"created_at"`
	UpdatedAt    time.Time    `db:"updated_at"`
	DeletedAt    sql.NullTime `db:"deleted_at"`
}

type UserDTO struct {
	Email       string    `json:"email"`
	Name        string    `json:"name"`
	Username    string    `json:"username"`
	Displayname string    `json:"displayName"`
	Bio         string    `json:"bio"`
}
