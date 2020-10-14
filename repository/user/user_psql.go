package userRepository

import (
	"api_backend_app/model"
	"database/sql"
)

// UserRepository : All postgresql related table interaction code lines will be here.
type UserRepository struct{}

// Register : postgresql db interaction for Register endpoint.
func (u UserRepository) Register(db *sql.DB, user model.User) (model.User, error) {
	stmt := "insert into users (email, password) values ($1,$2) RETURNING ID;"
	err := db.QueryRow(stmt, user.Email, user.Password).Scan(&user.ID)

	return user, err

}

// Login : postgresql db interaction for Login endpoint.
func (u UserRepository) Login(db *sql.DB, user model.User) (model.User, error) {
	stmt := "select * from users where email=$1"
	err := db.QueryRow(stmt, user.Email).Scan(&user.ID, &user.Email, &user.Password)

	return user, err
}
