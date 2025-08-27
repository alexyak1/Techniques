package controllers

import (
	"encoding/json"
	"main/database"
	"main/entity"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllKataTechniques(w http.ResponseWriter, r *http.Request) {
	kataTechniques := []entity.KataTechnique{}

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
	if kataName := query.Get("name"); kataName != "" {
		dbq = dbq.Where("kata_name = ?", kataName)
	}
	if serieName := query.Get("type"); serieName != "" {
		dbq = dbq.Where("type = ?", serieName)
	}
	dbq.Find(&kataTechniques)

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
	defer r.Body.Close()
	kataTechnique := entity.KataTechnique{}
	if err := json.NewDecoder(r.Body).Decode(&kataTechnique); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := database.Connector.Create(&kataTechnique).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

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
	defer r.Body.Close()
	kataTechnique := entity.KataTechnique{}

	if err := json.NewDecoder(r.Body).Decode(&kataTechnique); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	database.Connector.Save(&kataTechnique)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(kataTechnique)
}
