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
	"golang.org/x/crypto/bcrypt"
)

type Controller struct{}

func (c Controller) SignUp(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("signup invoked!")
		var user models.User
		var error models.Error

		json.NewDecoder(r.Body).Decode(&user)

		if user.Email == "" {
			error.Message = "Email is missing"
			utils.RespondWithError(w, http.StatusBadRequest, error)

			return
		}

		if user.Password == "" {
			error.Message = "Password is missing"
			utils.RespondWithError(w, http.StatusBadRequest, error)

			return
		}

		user.ClientType = "basic"

		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

		if err != nil {
			log.Fatal(err)
		}

		user.Password = string(hash)

		userRepo := userRepository.UserRepository{}
		user = userRepo.SignUp(db, user)

		if err != nil {
			error.Message = "Server error."
			utils.RespondWithError(w, http.StatusInternalServerError, error)
			return
		}

		user.Password = ""
		w.Header().Set("Content-Type", "application/json")
		utils.ResponseJSON(w, user)
	}
}
