package router

import (
	"net/http"

	"auth-jwt-golang/internal/auth"
	"auth-jwt-golang/internal/middleware"
	"auth-jwt-golang/internal/user"
)

// NewRouter menginisialisasi semua route aplikasi.
func NewRouter(
	authHandler *auth.Handler,
	userHandler *user.Handler,
	authService auth.ServiceToken,
	userService auth.Service,
) *http.ServeMux {
	mux := http.NewServeMux()

	// Public Routes
	mux.HandleFunc("POST /register", authHandler.Register)
	mux.HandleFunc("POST /login", authHandler.Login)

	// Protected Routes
	// Kita bungkus handler dengan middleware auth
	// Note: Go 1.22+ mendukung method di pattern (POST /login).
	// Untuk middleware, kita buat helper function atau wrapper.
	
	// Middleware wrapper
	authMiddleware := middleware.AuthMiddleware(authService, userService)

	// Route /users/profile dilindungi middleware
	mux.Handle("GET /users/profile", authMiddleware(http.HandlerFunc(userHandler.GetProfile)))

	return mux
}
