package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/eyupfatihersoy/app-tryout-1/models"
	userRepository "github.com/eyupfatihersoy/app-tryout-1/repository"
	"github.com/eyupfatihersoy/app-tryout-1/utils"
)

func (c Controller) ChangeUserType(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("user-type-changer service invoked!")
		var changedUserType models.ChangeUserType

		var clientType models.ChangeUserTypeRequestBody

		userEmail := r.Context().Value(models.ContextKey)
		json.NewDecoder(r.Body).Decode(&clientType)

		if validUserEmail, ok := userEmail.(models.EmailType); ok {
			changedUserType = models.ChangeUserType{ClientType: clientType.ClientType, Email: validUserEmail}
			//_ = changeUserType
		} else {
			return
		}

		userRepo := userRepository.UserRepository{}
		changedType, err := userRepo.ChangeUserType(db, changedUserType)

		if err != nil {
			log.Fatal(err)
		} else {
			utils.ResponseJSON(w, changedType)
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
