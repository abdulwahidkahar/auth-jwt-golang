package user

// UserResponse adalah format response untuk data user.
type UserResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
