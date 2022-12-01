package controllers

import (
	"encoding/json"
	"main/database"
	"main/entity"
	"net/http"
)

func GetBlogData(w http.ResponseWriter, r *http.Request) {
	var blogItems []entity.BlogItem

	database.Connector.Find(&blogItems)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(blogItems)
}
