package auth

import (
	"database/sql"
	"errors"
)

// Repository mendefinisikan kontrak/interface untuk akses data user.
// Interface ini memudahkan unit testing dengan mock.
type Repository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
}

// repository adalah implementasi dari interface Repository.
type repository struct {
	db *sql.DB
}

// NewRepository membuat instance baru dari repository.
func NewRepository(db *sql.DB) *repository {
	return &repository{db}
}

// Save menyimpan user baru ke database.
func (r *repository) Save(user User) (User, error) {
	// Query SQL untuk insert user
	query := "INSERT INTO users (name, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"
	
	// Eksekusi query
	result, err := r.db.Exec(query, user.Name, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return user, err
	}

	// Ambil ID yang baru saja di-generate
	id, err := result.LastInsertId()
	if err != nil {
		return user, err
	}

	user.ID = int(id)
	return user, nil
}

// FindByEmail mencari user berdasarkan email.
func (r *repository) FindByEmail(email string) (User, error) {
	var user User
	// Query SQL untuk select user
	query := "SELECT id, name, email, password, created_at, updated_at FROM users WHERE email = ?"

	// Eksekusi query dan mapping hasil ke struct User
	row := r.db.QueryRow(query, email)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			// Jika tidak ada data, kembalikan error khusus atau nil user
			return user, errors.New("user not found")
		}
		return user, err
	}

	return user, nil
}
