package user

import (
	"fmt"

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
	FindUserByEmail(string) (User, error)
}

func (r *userRepository) CreateUser(us User) error {
	fmt.Println("Chegou no user Repository")
	_, err := r.db.Db.Exec(insertUserQuery, us.Id, us.Name, us.Email, us.Password, us.Username, us.DisplayName, us.Bio, us.Birthday)
	
	return err	
}

func (r *userRepository) FindUserByEmail(email string) (User, error){
	var u User

	err := r.db.Db.Get(&u, findUserByEmailQuery, email)
	
	if err != nil {
		return u, err
	}
	
	return u, nil
}