package validator

import (
	"regexp"
)

// ValidateEmail memeriksa apakah format email valid menggunakan regex.
// Ini adalah validasi sederhana. Untuk validasi lebih kompleks, bisa gunakan library pihak ketiga.
func ValidateEmail(email string) bool {
	// Regex standar untuk email
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return re.MatchString(email)
}

// ValidatePassword memeriksa apakah password memenuhi kriteria keamanan minimal.
// Contoh: minimal 6 karakter.
func ValidatePassword(password string) bool {
	return len(password) >= 6
}
