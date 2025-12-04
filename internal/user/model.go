package user

import (
	"time"
)

// User struct di package user.
// Idealnya struct ini sama dengan di auth, atau auth mengimport ini.
// Karena struktur folder memisahkan, kita definisikan ulang atau import.
// Untuk kemandirian modul user, kita definisikan disini.
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
