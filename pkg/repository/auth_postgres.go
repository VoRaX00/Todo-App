package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"todoApp/Entity"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user Entity.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) VALUES ($1, $2, $3) RETURNING id", userTable)

	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return -1, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(login Entity.Login) (Entity.User, error) {
	var user Entity.User
	query := fmt.Sprintf("SELECT id,name FROM %s WHERE username=$1 AND password_hash=$2", userTable)
	err := r.db.Get(&user, query, login.Username, login.Password)

	return user, err
}
