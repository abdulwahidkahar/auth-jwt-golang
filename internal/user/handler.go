package user

import (
	"net/http"

	"auth-jwt-golang/internal/pkg/response"
)

type Handler struct {
	userService Service
}

func NewHandler(userService Service) *Handler {
	return &Handler{userService}
}

// GetProfile menangani request GET /users/profile.
// Endpoint ini dilindungi oleh middleware auth.
func (h *Handler) GetProfile(w http.ResponseWriter, r *http.Request) {
	// Ambil user ID dari context (diset oleh middleware)
	userID, ok := r.Context().Value("currentUser").(int)
	if !ok {
		response.JSON(w, http.StatusUnauthorized, response.ErrorResponse("Unauthorized", http.StatusUnauthorized, nil))
		return
	}

	// Ambil data user dari service
	user, err := h.userService.GetUserByID(userID)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, response.ErrorResponse(err.Error(), http.StatusBadRequest, nil))
		return
	}

	// Format response
	formatter := UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	response.JSON(w, http.StatusOK, response.SuccessResponse("User profile", http.StatusOK, formatter))
}
