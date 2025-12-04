package auth

import (
	"encoding/json"
	"net/http"

	"auth-jwt-golang/internal/pkg/response"
	"auth-jwt-golang/internal/pkg/validator"
)

// Handler struct menangani request HTTP untuk modul auth.
type Handler struct {
	authService Service
}

// NewHandler membuat instance baru dari Handler.
func NewHandler(authService Service) *Handler {
	return &Handler{authService}
}

// Register menangani endpoint POST /register.
func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var input RegisterRequest

	// Decode JSON body ke struct
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, response.ErrorResponse("Invalid request body", http.StatusBadRequest, nil))
		return
	}

	// Validasi input sederhana
	if input.Name == "" || input.Email == "" || input.Password == "" {
		response.JSON(w, http.StatusBadRequest, response.ErrorResponse("All fields are required", http.StatusBadRequest, nil))
		return
	}

	if !validator.ValidateEmail(input.Email) {
		response.JSON(w, http.StatusBadRequest, response.ErrorResponse("Invalid email format", http.StatusBadRequest, nil))
		return
	}

	if !validator.ValidatePassword(input.Password) {
		response.JSON(w, http.StatusBadRequest, response.ErrorResponse("Password must be at least 6 characters", http.StatusBadRequest, nil))
		return
	}

	// Panggil service untuk register
	newUser, err := h.authService.RegisterUser(input)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, response.ErrorResponse(err.Error(), http.StatusBadRequest, nil))
		return
	}

	// Format response
	formatter := UserResponse{
		ID:    newUser.ID,
		Name:  newUser.Name,
		Email: newUser.Email,
	}

	response.JSON(w, http.StatusOK, response.SuccessResponse("Account has been registered", http.StatusOK, formatter))
}

// Login menangani endpoint POST /login.
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var input LoginRequest

	// Decode JSON body
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, response.ErrorResponse("Invalid request body", http.StatusBadRequest, nil))
		return
	}

	// Validasi input
	if input.Email == "" || input.Password == "" {
		response.JSON(w, http.StatusBadRequest, response.ErrorResponse("Email and password are required", http.StatusBadRequest, nil))
		return
	}

	// Panggil service untuk login
	loggedinUser, token, err := h.authService.Login(input)
	if err != nil {
		response.JSON(w, http.StatusUnauthorized, response.ErrorResponse("Login failed", http.StatusUnauthorized, nil))
		return
	}

	// Format response dengan token
	formatter := UserResponse{
		ID:    loggedinUser.ID,
		Name:  loggedinUser.Name,
		Email: loggedinUser.Email,
		Token: token,
	}

	response.JSON(w, http.StatusOK, response.SuccessResponse("Login success", http.StatusOK, formatter))
}
