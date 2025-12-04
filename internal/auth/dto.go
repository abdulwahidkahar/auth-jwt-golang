package auth

// RegisterRequest adalah struktur data untuk input registrasi user.
type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginRequest adalah struktur data untuk input login user.
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserResponse adalah struktur data untuk response user yang aman (tanpa password).
type UserResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token,omitempty"` // Token JWT, opsional (hanya saat login/register)
}
