package response

import (
	"encoding/json"
	"net/http"
)

// Response adalah struktur standar untuk response API.
// Format ini konsisten untuk sukses maupun error.
type Response struct {
	Meta Meta        `json:"meta"` // Metadata response (status, message)
	Data interface{} `json:"data"` // Data payload (bisa null jika error)
}

// Meta menyimpan informasi status response.
type Meta struct {
	Message string `json:"message"` // Pesan deskriptif
	Code    int    `json:"code"`    // HTTP status code
	Status  string `json:"status"`  // "success" atau "error"
}

// SuccessResponse membuat struktur response sukses.
func SuccessResponse(message string, code int, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  "success",
	}

	return Response{
		Meta: meta,
		Data: data,
	}
}

// ErrorResponse membuat struktur response error.
func ErrorResponse(message string, code int, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  "error",
	}

	return Response{
		Meta: meta,
		Data: data,
	}
}

// JSON mengirimkan response dalam format JSON ke writer.
func JSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}
