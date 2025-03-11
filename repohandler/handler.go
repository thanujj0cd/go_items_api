package repohandler

import (
	"encoding/json"
	"groceriesApi/model"
	"net/http"
)

func GetGroceries(w http.ResponseWriter, r *http.Request) {
	db, err := ConnectDB()
	if err != nil {
		http.Error(w, "Failed to connect to the database", http.StatusInternalServerError)
		return
	}

	rows, err := db.Query("Select id, name, quantity FROM groceries")
	if err != nil {
		http.Error(w, "Failed to fetch groceries", http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var groceries []model.Grocery

	for rows.Next() {
		var g model.Grocery
		if err := rows.Scan(&g.ID, &g.Name, &g.Quantity); err != nil {
			http.Error(w, "Error scanning rows", http.StatusInternalServerError)
			return
		}
		groceries = append(groceries, g)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(groceries)
}
