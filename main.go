package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	controller "github.com/eyupfatihersoy/app-tryout-1/controllers"
	"github.com/eyupfatihersoy/app-tryout-1/driver"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

var db *sql.DB

func init() {
	gotenv.Load()
}

func main() {

	db = driver.ConnectDB()
	controllerIns := controller.Controller{}

	r := mux.NewRouter()
	r.HandleFunc("/signup", controllerIns.SignUp(db)).Methods("POST")
	r.HandleFunc("/login", controllerIns.LogIn(db)).Methods("POST")
	r.HandleFunc("/protected", controllerIns.TokenVerifyMiddleWare(controllerIns.ProtectedEndPoint())).Methods("GET")
	r.HandleFunc("/changeutype", controllerIns.TokenVerifyMiddleWare(controllerIns.ChangeUserType(db))).Methods("PUT")
	r.HandleFunc("/addtofavorites/{id}", controllerIns.AddToFavorites(db)).Methods("POST")

	fmt.Println("hello")

	log.Fatal(http.ListenAndServe(":8000", r))
}
