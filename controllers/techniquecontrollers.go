package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"main/database"
	"main/entity"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllTechniques(w http.ResponseWriter, r *http.Request) {
	var techniques []entity.Technique

	if belt, ok := r.URL.Query()["belt"]; ok {
		database.Connector.Where("belt = ?", belt).Find(&techniques)
	} else {
		database.Connector.Find(&techniques)
	}

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
	requestBody, _ := ioutil.ReadAll(r.Body)
	var technique entity.Technique
	json.Unmarshal(requestBody, &technique)

	query := fmt.Sprintf(
		"INSERT INTO techniques (name, belt, image_url, type) "+
			"VALUES ('%s', '%s', '%s', '%s');",
		technique.Name, technique.Belt, technique.ImageURL, technique.Type,
	)

	database.Connector.DB().QueryRow(query)
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
	requestBody, _ := ioutil.ReadAll(r.Body)
	var technique entity.Technique
	json.Unmarshal(requestBody, &technique)
	database.Connector.Save(&technique)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(technique)
}
