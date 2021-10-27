package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/eyupfatihersoy/app-tryout-1/models"
)

func (c Controller) ChangeUserType(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("user-type-changer service invoked!")

		var clientType models.ClientType
		_ = clientType

		json.NewDecoder(r.Body).Decode(&clientType)
	}

}

func (c Controller) ProtectedEndPoint() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("protected endpoint invoked!")
		test := r.Context().Value(models.ContextKey)
		fmt.Println(test)
	}
}
