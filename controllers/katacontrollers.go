package controllers

import (
	"encoding/json"
	"io/ioutil"
	"main/database"
	"main/entity"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetAllKataTechniques(w http.ResponseWriter, r *http.Request) {
	kataTechniques := []entity.KataTechnique{}
	var err error

	if kataName, ok := r.URL.Query()["name"]; ok {
		err = database.Connector.Where("kata_name = ?", kataName).Find(&kataTechniques).Error
	} else if serie_name, ok := r.URL.Query()["type"]; ok {
		err = database.Connector.Where("type = ?", serie_name).Find(&kataTechniques).Error
	} else {
		err = database.Connector.Find(&kataTechniques).Error
	}

	if err != nil {
		http.Error(w, "Failed to retrieve kata techniques: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(kataTechniques); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func GetKataTechniqueById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	if key == "" {
		http.Error(w, "Kata technique ID is required", http.StatusBadRequest)
		return
	}

	kataTechnique := entity.KataTechnique{}
	err := database.Connector.First(&kataTechnique, key).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Kata technique not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to retrieve kata technique: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(kataTechnique); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func CreateKataTechnique(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	kataTechnique := entity.KataTechnique{}
	if err := json.Unmarshal(requestBody, &kataTechnique); err != nil {
		http.Error(w, "Invalid JSON format: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Validate required fields
	if kataTechnique.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}
	if kataTechnique.KataName == "" {
		http.Error(w, "Kata name is required", http.StatusBadRequest)
		return
	}

	// Use GORM Create instead of raw SQL to prevent injection
	if err := database.Connector.Create(&kataTechnique).Error; err != nil {
		http.Error(w, "Failed to create kata technique: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(kataTechnique); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func DeleteKataTechniqueById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	if key == "" {
		http.Error(w, "Kata technique ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(key, 10, 64)
	if err != nil {
		http.Error(w, "Invalid kata technique ID format", http.StatusBadRequest)
		return
	}

	kataTechnique := entity.KataTechnique{}
	result := database.Connector.Where("id = ?", id).Delete(&kataTechnique)
	if err := result.Error; err != nil {
		http.Error(w, "Failed to delete kata technique: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0 {
		http.Error(w, "Kata technique not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func UpdateKataTechniqueById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	if key == "" {
		http.Error(w, "Kata technique ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(key, 10, 64)
	if err != nil {
		http.Error(w, "Invalid kata technique ID format", http.StatusBadRequest)
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	kataTechnique := entity.KataTechnique{}
	if err := json.Unmarshal(requestBody, &kataTechnique); err != nil {
		http.Error(w, "Invalid JSON format: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Set the ID from URL parameter
	kataTechnique.Id = &key

	// Check if kata technique exists
	var existingKataTechnique entity.KataTechnique
	if err := database.Connector.First(&existingKataTechnique, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Kata technique not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to find kata technique: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Update the kata technique
	if err := database.Connector.Save(&kataTechnique).Error; err != nil {
		http.Error(w, "Failed to update kata technique: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(kataTechnique); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
