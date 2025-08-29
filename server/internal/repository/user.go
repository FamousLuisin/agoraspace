package repository

import (
	"time"

	"github.com/FamousLuisin/agoraspace/internal/db"
	"github.com/FamousLuisin/agoraspace/internal/models"
	"github.com/jmoiron/sqlx"
)

const (
	insertUserQuery = `
		INSERT INTO tb_users (
			id,
			name,
			email,
			password,
			username,
			display_name,
			bio,
			birthday
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7 ,$8
		)
	`
	findUserById = `
		SELECT * FROM tb_users WHERE id = $1
	`
	findUserByEmailQuery = `
		SELECT * FROM tb_users WHERE email = $1
	`
	findUserByUsernameQuery = `
		SELECT * FROM tb_users WHERE username = $1 AND deleted_at IS NULL
	`
	getAllUsersQuery = `
		SELECT * FROM tb_users WHERE deleted_at IS NULL OFFSET $1 LIMIT $2
	`
	updateUserQuery = `
		UPDATE tb_users 
		SET email = $1, name = $2, username = $3, display_name = $4, bio = $5, updated_at = $6
		WHERE id = $7
	`
	deleteUserQuery = `
		UPDATE tb_users
		SET deleted_at = $1
		WHERE id = $2
	`
)

func NewUserRepository(db *db.Database) UserRepository {
	return &userRepository{
		db: db.Db,
	}
}

type userRepository struct{
	db *sqlx.DB
}

type UserRepository interface {
	CreateUser(models.User) error
	FindUserById(string) (*models.User, error)
	FindUserByEmail(string) (*models.User, error)
	FindUserByUsername(string) (*models.User, error)
	GetAllUsers(int, int) (*[]models.User, error)
	UpdateUser(models.User) error
	DeleteUser(models.User) error
	ActivateUser(models.User) error
}

func (r *userRepository) CreateUser(us models.User) error {
	_, err := r.db.Exec(insertUserQuery, us.Id, us.Name, us.Email, us.Password, us.Username, us.DisplayName, us.Bio, us.Birthday)
	
	return err	
}

func (r *userRepository) FindUserByEmail(email string) (*models.User, error){
	var u models.User

	err := r.db.Get(&u, findUserByEmailQuery, email)
	
	if err != nil {
		return nil, err
	}
	
	return &u, nil
}

func (r *userRepository) FindUserByUsername(username string) (*models.User, error){
	var u models.User

	if err := r.db.Get(&u, findUserByUsernameQuery, username); err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *userRepository) GetAllUsers(page, perPage int) (*[]models.User, error){
	var users []models.User

	if err := r.db.Select(&users, getAllUsersQuery, page, perPage); err != nil {
		return nil, err
	}

	return &users, nil
}

func (r *userRepository) UpdateUser(u models.User) error {
	_, err := r.db.Exec(updateUserQuery, u.Email, u.Name, u.Username, u.DisplayName, u.Bio, time.Now(), u.Id)

	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) DeleteUser(u models.User) error {
	_, err := r.db.Exec(deleteUserQuery, time.Now(), u.Id)
	
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) ActivateUser(u models.User) error {
	_, err := r.db.Exec(deleteUserQuery, nil, u.Id)
	
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) FindUserById(identifier string) (*models.User, error){
	var u models.User

	err := r.db.Get(&u, findUserById, identifier)

	if err != nil {
		return nil, err
	}

	return &u, nil
}