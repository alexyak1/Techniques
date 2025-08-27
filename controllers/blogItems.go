package controllers

import (
	"encoding/json"
	"main/database"
	"main/entity"
	"net/http"
	"strconv"
)

func GetBlogData(w http.ResponseWriter, r *http.Request) {
	var blogItems []entity.BlogItem

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

	database.Connector.Limit(limit).Offset(offset).Find(&blogItems)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(blogItems)
}

func GetBlogDataRawQuery(w http.ResponseWriter, r *http.Request) {
	var blogItems []entity.BlogItem

	// With join to pool answers
	database.Connector.Joins("JOIN blog_pool_answers on blog_pool_answers.blog_item_id = blog_items.id").Find(&blogItems)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(blogItems)
}
