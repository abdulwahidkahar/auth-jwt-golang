package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"auth-jwt-golang/internal/config"

	_ "github.com/go-sql-driver/mysql" // Import driver MySQL secara blank identifier agar init() driver dijalankan
)

// DB adalah variabel global yang menyimpan koneksi database pool.
// Digunakan di seluruh aplikasi untuk melakukan query ke database.
var DB *sql.DB

// InitDB menginisialisasi koneksi ke database MySQL.
// Fungsi ini membaca konfigurasi dari package config dan membuka koneksi.
func InitDB() {
	// Format DSN (Data Source Name) untuk MySQL: user:password@tcp(host:port)/dbname?parseTime=true
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.Config.DBUser,
		config.Config.DBPassword,
		config.Config.DBHost,
		config.Config.DBPort,
		config.Config.DBName,
	)

	var err error
	// sql.Open tidak langsung membuat koneksi, hanya menyiapkan handle.
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Gagal membuka koneksi database: %v", err)
	}

	// Ping database untuk memastikan koneksi benar-benar terbentuk.
	if err = DB.Ping(); err != nil {
		log.Fatalf("Gagal terhubung ke database: %v", err)
	}

	// Konfigurasi Connection Pool
	// SetMaxOpenConns: Jumlah maksimal koneksi yang terbuka ke database.
	DB.SetMaxOpenConns(25)
	// SetMaxIdleConns: Jumlah maksimal koneksi idle (menganggur) yang disimpan di pool.
	DB.SetMaxIdleConns(25)
	// SetConnMaxLifetime: Lama waktu maksimal sebuah koneksi boleh digunakan sebelum didaur ulang.
	DB.SetConnMaxLifetime(5 * time.Minute)

	log.Println("Berhasil terhubung ke database!")
}