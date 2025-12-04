package hash

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword mengubah password plain text menjadi hash menggunakan algoritma bcrypt.
// Bcrypt dipilih karena aman dan lambat (computationally expensive), sehingga sulit di-brute force.
func HashPassword(password string) (string, error) {
	// GenerateFromPassword melakukan hashing.
	// Cost adalah tingkat kompleksitas hashing. DefaultCost (10) adalah standar yang baik.
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash membandingkan password plain text dengan hash yang tersimpan.
// Mengembalikan true jika cocok, false jika tidak.
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
