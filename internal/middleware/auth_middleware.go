package middleware

import (
	"context"
	"net/http"
	"strings"

	"auth-jwt-golang/internal/auth"
	"auth-jwt-golang/internal/pkg/response"

	"github.com/golang-jwt/jwt/v5"
)

// AuthMiddleware memvalidasi JWT token di header Authorization.
// Jika valid, user_id akan disimpan di context request.
func AuthMiddleware(authService auth.ServiceToken, userService auth.Service) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// 1. Ambil header Authorization
			authHeader := r.Header.Get("Authorization")
			if !strings.Contains(authHeader, "Bearer") {
				response.JSON(w, http.StatusUnauthorized, response.ErrorResponse("Unauthorized", http.StatusUnauthorized, nil))
				return
			}

			// 2. Ambil token string (pisahkan "Bearer " dari token)
			tokenString := ""
			arrayToken := strings.Split(authHeader, " ")
			if len(arrayToken) == 2 {
				tokenString = arrayToken[1]
			}

			// 3. Validasi token
			token, err := authService.ValidateToken(tokenString)
			if err != nil {
				response.JSON(w, http.StatusUnauthorized, response.ErrorResponse("Invalid Token", http.StatusUnauthorized, nil))
				return
			}

			// 4. Ambil claims dari token
			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok || !token.Valid {
				response.JSON(w, http.StatusUnauthorized, response.ErrorResponse("Unauthorized", http.StatusUnauthorized, nil))
				return
			}

			// 5. Ambil User ID dari claims
			userID := int(claims["user_id"].(float64))

			// 6. Simpan User ID ke context (bisa juga load full user dari DB jika perlu)
			// Disini kita simpan userID saja agar ringan, atau bisa simpan struct User jika query DB.
			// Untuk best practice performa, simpan ID saja cukup, tapi jika butuh data user di controller, query DB disini.
			// Kita simpan ID nya saja di context dengan key "currentUser"
			ctx := context.WithValue(r.Context(), "currentUser", userID)

			// Lanjut ke handler berikutnya dengan context baru
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
