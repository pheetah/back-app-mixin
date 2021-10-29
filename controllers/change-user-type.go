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
		var changeUserType models.ChangeUserType

		var clientType models.ChangeUserTypeBody

		userEmail := r.Context().Value(models.ContextKey)
		json.NewDecoder(r.Body).Decode(&clientType)

		if validUserEmail, ok := userEmail.(models.EmailType); ok {
			changeUserType = models.ChangeUserType{clientType.ClientType, validUserEmail}
			_ = changeUserType
		} else {
			return
		}

	}

}

func (c Controller) ProtectedEndPoint() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("protected endpoint invoked!")
		test := r.Context().Value(models.ContextKey)
		fmt.Println(test)
	}
}
