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

func GetAllKataTechniques(w http.ResponseWriter, r *http.Request) {
	kataTechniques := []entity.KataTechnique{}

	if kataName, ok := r.URL.Query()["name"]; ok {
		database.Connector.Where("kata_name = ?", kataName).Find(&kataTechniques)
	} else if serie_name, ok := r.URL.Query()["type"]; ok {
		database.Connector.Where("type = ?", serie_name).Find(&kataTechniques)
	} else {
		database.Connector.Find(&kataTechniques)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(kataTechniques)
}

func GetKataTechniqueById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	kataTechnique := entity.KataTechnique{}
	database.Connector.First(&kataTechnique, key)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(kataTechnique)
}

func CreateKataTechnique(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	kataTechnique := entity.KataTechnique{}
	json.Unmarshal(requestBody, &kataTechnique)

	query := fmt.Sprintf(
		"INSERT INTO kata_techniques (name, kata_name, image_url, type, image_id) "+
			"VALUES ('%s', '%s', '%s', '%s', '%s');",
		kataTechnique.Name, kataTechnique.KataName, kataTechnique.ImageURL, kataTechnique.Type, kataTechnique.ImageId,
	)

	database.Connector.DB().QueryRow(query)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(kataTechnique)
}

func DeleteKataTechniqueById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	kataTechnique := entity.KataTechnique{}
	id, _ := strconv.ParseInt(key, 10, 64)
	database.Connector.Where("id = ?", id).Delete(&kataTechnique)
	w.WriteHeader(http.StatusNoContent)
}

func UpdateKataTechniqueById(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	kataTechnique := entity.KataTechnique{}

	json.Unmarshal(requestBody, &kataTechnique)
	database.Connector.Save(&kataTechnique)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(kataTechnique)
}
