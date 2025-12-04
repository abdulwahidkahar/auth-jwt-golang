package auth

import (
	"errors"

	"auth-jwt-golang/internal/pkg/hash"
	"auth-jwt-golang/internal/utils"
)

// Service mendefinisikan kontrak untuk business logic auth.
type Service interface {
	RegisterUser(input RegisterRequest) (User, error)
	Login(input LoginRequest) (User, string, error)
}

type service struct {
	repository Repository
	jwtService ServiceToken
}

// NewService membuat instance baru dari service auth.
func NewService(repository Repository, jwtService ServiceToken) *service {
	return &service{repository, jwtService}
}

// RegisterUser menangani proses pendaftaran user baru.
func (s *service) RegisterUser(input RegisterRequest) (User, error) {
	// 1. Cek apakah email sudah terdaftar
	existingUser, err := s.repository.FindByEmail(input.Email)
	if err == nil && existingUser.ID != 0 {
		return User{}, errors.New("email already registered")
	}

	// 2. Hash password
	passwordHash, err := hash.HashPassword(input.Password)
	if err != nil {
		return User{}, err
	}

	// 3. Buat object user baru
	user := User{
		Name:      input.Name,
		Email:     input.Email,
		Password:  passwordHash,
		CreatedAt: utils.GetCurrentTime(),
		UpdatedAt: utils.GetCurrentTime(),
	}

	// 4. Simpan ke database
	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

// Login menangani proses login user.
func (s *service) Login(input LoginRequest) (User, string, error) {
	email := input.Email
	password := input.Password

	// 1. Cari user berdasarkan email
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, "", errors.New("invalid email or password")
	}

	// 2. Verifikasi password
	if !hash.CheckPasswordHash(password, user.Password) {
		return user, "", errors.New("invalid email or password")
	}

	// 3. Generate JWT token
	token, err := s.jwtService.GenerateToken(user.ID)
	if err != nil {
		return user, "", err
	}

	return user, token, nil
}
