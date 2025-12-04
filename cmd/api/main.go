package main

import (
	"fmt"
	"log"
	"net/http"

	"auth-jwt-golang/internal/auth"
	"auth-jwt-golang/internal/config"
	"auth-jwt-golang/internal/database"
	"auth-jwt-golang/internal/router"
	"auth-jwt-golang/internal/user"
)

func main() {
	// 1. Load Konfigurasi
	config.LoadConfig()

	// 2. Init Database
	database.InitDB()
	// Defer close database connection saat aplikasi berhenti
	defer database.DB.Close()

	// 3. Init Repository, Service, dan Handler (Dependency Injection)
	
	// Auth Module
	authRepo := auth.NewRepository(database.DB)
	tokenService := auth.NewServiceToken()
	authService := auth.NewService(authRepo, tokenService)
	authHandler := auth.NewHandler(authService)

	// User Module
	userRepo := user.NewRepository(database.DB)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

	// 4. Init Router
	r := router.NewRouter(authHandler, userHandler, tokenService, authService)

	// 5. Jalankan Server
	// Default port 8080
	serverPort := ":8080"

	fmt.Printf("Server berjalan di http://localhost%s\n", serverPort)
	err := http.ListenAndServe(serverPort, r)
	if err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}