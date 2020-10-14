package main

import (
	"api_backend_app/controller"
	"api_backend_app/driver"
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

var db *sql.DB

func main() {
	// loads Environment variables in .env file
	gotenv.Load()

	db = driver.ConnectDB()
	ctrlr := controller.Controller{}

	// Create the router & register API Endpoints and it's handler functions.
	router := mux.NewRouter()
	router.HandleFunc("/register", ctrlr.Register(db)).Methods("POST")
	router.HandleFunc("/login", ctrlr.Login(db)).Methods("POST")
	router.HandleFunc("/protected", ctrlr.TokenVerifyMiddleware(ctrlr.ProtectedEndpoint())).Methods("GET")

	// Start HTTP server
	log.Println("Listen on the port 8000...")
	log.Fatal(http.ListenAndServe(":8000", router)) // log.Fatal terminates the program (after printing the inside content) ONLY IF error is non-nil.
}
