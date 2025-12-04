package auth

import (
	"time"
)

// User merepresentasikan tabel 'users' di database.
// Struct ini digunakan oleh repository untuk mapping data dari database ke object Go.
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"` // Password tidak akan dirender ke JSON response (keamanan)
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
