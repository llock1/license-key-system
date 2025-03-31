package handlers

import (
	"encoding/json"
	"fmt"
	"license/database"
	"license/models"
	"net/http"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func AddKey(w http.ResponseWriter, req *http.Request) {

	//err := auth.AuthenticateRequest(req)
	//if err != nil {
	//	w.WriteHeader(http.StatusUnauthorized)
	//	return
	//}

	type Response struct {
		Message string `json:"message"`
		Success bool   `json:"success"`
		Key     string `json:"key"`
	}

	new_uuid := uuid.New().String()
	key := models.LicenseKey{Key: new_uuid}
	database.Client.Create(&key)
	response := Response{
		Message: "Successfullt Created Key",
		Success: true,
		Key:     new_uuid,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func DeleteKey(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	keyUUID := req.URL.Query().Get("uuid")
	if keyUUID == "" {
		http.Error(w, "Missing UUID", http.StatusBadRequest)
		return
	}

	var key models.LicenseKey
	err := database.Client.Where("key = ?", keyUUID).First(&key).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Key does not exist", http.StatusNotFound)
		} else {
			http.Error(w, "Database error", http.StatusInternalServerError)
		}
		return
	}

	err = database.Client.Delete(&key).Error
	if err != nil {
		http.Error(w, "Failed to delete key", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Deleted Key: %s", key.Key)
}

func AuthenticateKey(w http.ResponseWriter, req *http.Request) {

	type Response struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
		Key     string `json:"key"`
	}

	if req.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	uuid := req.URL.Query().Get("uuid")
	var key models.LicenseKey
	err := database.Client.Where("key = ?", uuid).First(&key).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Key Not Found", http.StatusNotFound)
		} else {
			http.Error(w, "Database error", http.StatusInternalServerError)
		}
		return
	}

	response := Response{
		Success: true,
		Message: "Successfully authenticated",
		Key:     uuid,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
