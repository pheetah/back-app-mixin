package controller

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/eyupfatihersoy/app-tryout-1/models"
	userRepository "github.com/eyupfatihersoy/app-tryout-1/repository"
	"github.com/eyupfatihersoy/app-tryout-1/utils"
	"github.com/golang-jwt/jwt"
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

func (c Controller) LogIn(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("login invoked!")

		var user models.User
		var jwt models.JWT
		var error models.Error
		var response models.LoginResponse

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

		password := user.Password

		userRepo := userRepository.UserRepository{}
		user, err := userRepo.LogIn(db, user)

		if err != nil {
			if err == sql.ErrNoRows {
				error.Message = "The user does not exist"
				utils.RespondWithError(w, http.StatusBadRequest, error)
				return
			} else {
				log.Fatal(err)
			}
		}

		hashedPassword := user.Password

		err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

		if err != nil {
			error.Message = "Invalid Password"
			utils.RespondWithError(w, http.StatusUnauthorized, error)
			return
		}

		token, err := utils.GenerateToken(user)

		if err != nil {
			log.Fatal(err)
		}

		w.WriteHeader(http.StatusOK)
		jwt.Token = token

		response.JWT = jwt.Token
		response.ClientType = user.ClientType

		utils.ResponseJSON(w, response)
	}
}

func (c Controller) TokenVerifyMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	fmt.Println("Token Middleware invoked!")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var errorObject models.Error
		authHeader := r.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")
		var email models.EmailType

		if len(bearerToken) == 2 {
			authToken := bearerToken[1]

			token, error := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("there was an error")
				}

				return []byte(os.Getenv("SECRET")), nil
			})

			claims, _ := token.Claims.(jwt.MapClaims)

			for key, val := range claims {
				if key == "email" {
					if str, ok := val.(string); ok {
						/* act on str */
						email = models.EmailType(str)
						fmt.Println(email)
					} else {
						/* not string */
					}

				} else {
					_ = val
					_ = email
				}
			}

			if error != nil {
				errorObject.Message = error.Error()
				utils.RespondWithError(w, http.StatusUnauthorized, errorObject)
				return
			}

			if token.Valid {
				ctx := context.WithValue(r.Context(), models.ContextKey, email)
				next.ServeHTTP(w, r.WithContext(ctx))
			} else {
				errorObject.Message = error.Error()
				utils.RespondWithError(w, http.StatusUnauthorized, errorObject)
				return
			}

		} else {
			errorObject.Message = "invalid token."
			utils.RespondWithError(w, http.StatusUnauthorized, errorObject)
			return
		}
	})
}
