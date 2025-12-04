package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// AppConfig menyimpan konfigurasi aplikasi yang dimuat dari environment variables.
// Struct ini digunakan agar konfigurasi terpusat dan mudah diakses di seluruh aplikasi.
type AppConfig struct {
	AppName    string // Nama aplikasi
	JWTSecret  string // Secret key untuk menandatangani JWT token
	DBHost     string // Host database (misal: localhost atau IP)
	DBUser     string // Username database
	DBPassword string // Password database
	DBName     string // Nama database
	DBPort     string // Port database
}

// Config adalah variabel global yang menyimpan konfigurasi aplikasi.
// Variabel ini akan diisi saat fungsi LoadConfig dipanggil.
var Config AppConfig

// LoadConfig memuat konfigurasi dari file .env dan environment variables.
// Fungsi ini harus dipanggil di awal aplikasi (misal: di main.go).
func LoadConfig() {
	// Memuat file .env jika ada.
	// Mencoba memuat dari direktori saat ini, atau dari root project jika dijalankan dari cmd/api
	err := godotenv.Load()
	if err != nil {
		// Jika gagal, coba load dari parent directory (untuk kasus dijalankan dari cmd/api)
		_ = godotenv.Load("../../.env")
	}

	Config = AppConfig{
		AppName:    getEnv("APP_NAME", "MyApp"),
		JWTSecret:  getEnv("JWT_SECRET", "default_secret"), // Penting: Ganti ini di production!
		DBHost:     getEnv("DB_HOST", "127.0.0.1"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", "mydb"),
		DBPort:     getEnv("DB_PORT", "3306"),
	}
}

// getEnv mengambil nilai dari environment variable.
// Jika key tidak ditemukan, maka akan mengembalikan nilai defaultValue.
// Ini membantu aplikasi tetap berjalan dengan konfigurasi default jika env var lupa diset.
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	return value
}

// getEnvInt mengambil nilai integer dari environment variable.
// Berguna jika ada konfigurasi yang berupa angka (misal: port, timeout).
func getEnvInt(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}
	return value
}
