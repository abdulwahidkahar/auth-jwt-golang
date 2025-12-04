package auth

import (
	"errors"
	"time"

	"auth-jwt-golang/internal/config"

	"github.com/golang-jwt/jwt/v5"
)

// ServiceToken mendefinisikan kontrak untuk layanan token.
type ServiceToken interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
}

// NewServiceToken membuat instance baru dari jwtService.
func NewServiceToken() *jwtService {
	return &jwtService{}
}

// GenerateToken membuat JWT token baru untuk user.
func (s *jwtService) GenerateToken(userID int) (string, error) {
	// Membuat claim (payload) token
	claims := jwt.MapClaims{}
	claims["user_id"] = userID
	// Token berlaku selama 24 jam
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Membuat token dengan algoritma HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Menandatangani token dengan secret key
	signedToken, err := token.SignedString([]byte(config.Config.JWTSecret))
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

// ValidateToken memvalidasi apakah token valid dan mengembalikan objek token.
func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	// Parse token
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		// Validasi algoritma signing
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token signing method")
		}
		// Kembalikan secret key untuk verifikasi
		return []byte(config.Config.JWTSecret), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
