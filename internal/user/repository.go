package user

import (
	"database/sql"
	"errors"
)

type Repository interface {
	FindByID(ID int) (User, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{db}
}

func (r *repository) FindByID(ID int) (User, error) {
	var user User
	query := "SELECT id, name, email, password, created_at, updated_at FROM users WHERE id = ?"
	
	row := r.db.QueryRow(query, ID)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return user, errors.New("user not found")
		}
		return user, err
	}

	return user, nil
}
