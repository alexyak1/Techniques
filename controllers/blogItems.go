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

func GetBlogDataRawQuery(w http.ResponseWriter, r *http.Request) {
	var blogItems []entity.BlogItem

	// only blog_items
	// database.Connector.Find(&blogItems)

	// TODO: was bad connection need to test with join
	// With join to pool answers
	database.Connector.Joins("JOIN blog_pool_answers on blog_pool_answers.blog_item_id = blog_items.id").Find(&blogItems)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(blogItems)
}
