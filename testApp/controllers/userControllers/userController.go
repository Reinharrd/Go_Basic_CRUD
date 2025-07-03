package userControllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"testApp/models/userModels"
)

func GetDataUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 2
	}

	users, err := userModels.GetDataUser(page, limit)

	if err != nil {
		http.Error(w, "Gagal mengambil data user", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(users)

	if err != nil {
		http.Error(w, "Gagal encode JSON", http.StatusInternalServerError)
		return
	}

}
