package controllers

import (
	"encoding/json"
	"io/ioutil"
	"main/database"
	"main/entity"
	"net/http"
	"strconv"

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

func CreateTechnique(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var technique entity.Technique
	json.Unmarshal(requestBody, &technique)

	database.Connector.Create(technique)
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