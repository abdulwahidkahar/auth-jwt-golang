package utils

import "time"

// GetCurrentTime mengembalikan waktu saat ini dalam UTC.
// Menggunakan UTC disarankan untuk konsistensi di server yang berbeda zona waktu.
func GetCurrentTime() time.Time {
	return time.Now().UTC()
}

// FormatTime mengubah waktu menjadi string format standar (ISO 8601 / RFC 3339).
func FormatTime(t time.Time) string {
	return t.Format(time.RFC3339)
}
