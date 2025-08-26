package user

import (
	"fmt"
	"time"

	"github.com/FamousLuisin/agoraspace/internal/db"
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
		db: db,
	}
}

type userRepository struct{
	db *db.Database
}

type UserRepository interface {
	CreateUser(User) error
	FindUserByEmail(string) (*User, error)
	FindUserByUsername(string) (*User, error)
	GetAllUsers(int, int) (*[]User, error)
	UpdateUser(User) error
	DeleteUser(User) error
	ActivateUser(User) error
}

func (r *userRepository) CreateUser(us User) error {
	fmt.Println("Chegou no user Repository")
	_, err := r.db.Db.Exec(insertUserQuery, us.Id, us.Name, us.Email, us.Password, us.Username, us.DisplayName, us.Bio, us.Birthday)
	
	return err	
}

func (r *userRepository) FindUserByEmail(email string) (*User, error){
	var u User

	err := r.db.Db.Get(&u, findUserByEmailQuery, email)
	
	if err != nil {
		return nil, err
	}
	
	return &u, nil
}

func (r *userRepository) FindUserByUsername(username string) (*User, error){
	var u User

	if err := r.db.Db.Get(&u, findUserByUsernameQuery, username); err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *userRepository) GetAllUsers(page, perPage int) (*[]User, error){
	var users []User

	if err := r.db.Db.Select(&users, getAllUsersQuery, page, perPage); err != nil {
		return nil, err
	}

	return &users, nil
}

func (r *userRepository) UpdateUser(u User) error {
	_, err := r.db.Db.Exec(updateUserQuery, u.Email, u.Name, u.Username, u.DisplayName, u.Bio, time.Now(), u.Id)

	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) DeleteUser(u User) error {
	_, err := r.db.Db.Exec(deleteUserQuery, time.Now(), u.Id)
	
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) ActivateUser(u User) error {
	_, err := r.db.Db.Exec(deleteUserQuery, nil, u.Id)
	
	if err != nil {
		return err
	}

	return nil
}