package controllers

import (
	"encoding/json"
	"main/database"
	"main/entity"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllTechniques(w http.ResponseWriter, r *http.Request) {
	var techniques []entity.Technique

	query := r.URL.Query()
	limit := 50
	offset := 0
	if v := query.Get("limit"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 && n <= 200 {
			limit = n
		}
	}
	if v := query.Get("offset"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n >= 0 {
			offset = n
		}
	}

	dbq := database.Connector.Limit(limit).Offset(offset)
	if belt := query.Get("belt"); belt != "" {
		dbq = dbq.Where("belt = ?", belt)
	}
	dbq.Find(&techniques)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(techniques)
}

func GetTechniqueById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var technique entity.Technique
	database.Connector.First(&technique, key)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(technique)
}

func CreateTechnique(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var technique entity.Technique
	if err := json.NewDecoder(r.Body).Decode(&technique); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := database.Connector.Create(&technique).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(technique)
}

func DeleteTechniqueById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var technique entity.Technique
	id, _ := strconv.ParseInt(key, 10, 64)
	database.Connector.Where("id = ?", id).Delete(&technique)
	w.WriteHeader(http.StatusNoContent)
}

func UpdateTechniqueById(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var technique entity.Technique
	if err := json.NewDecoder(r.Body).Decode(&technique); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	database.Connector.Save(&technique)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(technique)
}
