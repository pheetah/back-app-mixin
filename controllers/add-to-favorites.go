package controller

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/eyupfatihersoy/app-tryout-1/utils"
	"github.com/gorilla/mux"
)

func (c Controller) AddToFavorites(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		fmt.Println("params invoked", params)

		utils.ResponseJSON(w, params)
	}

}
