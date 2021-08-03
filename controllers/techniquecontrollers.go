package controllers

import (
	"encoding/json"
	"main/database"
	"main/entity"
	"net/http"
)


func GetAllTechniques(w http.ResponseWriter, r *http.Request) {
	var techniques []entity.Technique
	database.Connector.Find(&techniques)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(techniques)
}