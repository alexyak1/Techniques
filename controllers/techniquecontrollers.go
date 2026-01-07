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

func GetAllTechniques(w http.ResponseWriter, r *http.Request) {
	var techniques []entity.Technique
	var err error

	if belt, ok := r.URL.Query()["belt"]; ok {
		err = database.Connector.Where("belt = ?", belt).Find(&techniques).Error
	} else {
		err = database.Connector.Find(&techniques).Error
	}

	if err != nil {
		http.Error(w, "Failed to retrieve techniques: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(techniques); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func GetTechniqueById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	if key == "" {
		http.Error(w, "Technique ID is required", http.StatusBadRequest)
		return
	}

	var technique entity.Technique
	err := database.Connector.First(&technique, key).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Technique not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to retrieve technique: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(technique); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func CreateTechnique(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var technique entity.Technique
	if err := json.Unmarshal(requestBody, &technique); err != nil {
		http.Error(w, "Invalid JSON format: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Validate required fields
	if technique.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}
	if technique.Belt == "" {
		http.Error(w, "Belt is required", http.StatusBadRequest)
		return
	}

	// Use GORM Create instead of raw SQL to prevent injection
	if err := database.Connector.Create(&technique).Error; err != nil {
		http.Error(w, "Failed to create technique: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(technique); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func DeleteTechniqueById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	if key == "" {
		http.Error(w, "Technique ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(key, 10, 64)
	if err != nil {
		http.Error(w, "Invalid technique ID format", http.StatusBadRequest)
		return
	}

	var technique entity.Technique
	result := database.Connector.Where("id = ?", id).Delete(&technique)
	if err := result.Error; err != nil {
		http.Error(w, "Failed to delete technique: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0 {
		http.Error(w, "Technique not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func UpdateTechniqueById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	if key == "" {
		http.Error(w, "Technique ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(key, 10, 64)
	if err != nil {
		http.Error(w, "Invalid technique ID format", http.StatusBadRequest)
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var technique entity.Technique
	if err := json.Unmarshal(requestBody, &technique); err != nil {
		http.Error(w, "Invalid JSON format: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Set the ID from URL parameter
	technique.Id = &key

	// Check if technique exists
	var existingTechnique entity.Technique
	if err := database.Connector.First(&existingTechnique, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Technique not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to find technique: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Update the technique
	if err := database.Connector.Save(&technique).Error; err != nil {
		http.Error(w, "Failed to update technique: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(technique); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
