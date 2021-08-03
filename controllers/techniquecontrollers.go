package controllers

import (
	"encoding/json"
	"main/database"
	"main/entity"
	"net/http"

	"github.com/gorilla/mux"
)


func GetAllTechniques(w http.ResponseWriter, r *http.Request) {
	var techniques []entity.Technique
	database.Connector.Find(&techniques)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(techniques)
}

func GetTechniqueById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    key := vars["id"]

	var technique entity.Technique
	database.Connector.First(&technique, key)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(technique)
}