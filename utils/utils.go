package utils

import (
	"api_backend_app/model"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// RespondWithError : to give error response
func RespondWithError(w http.ResponseWriter, status int, error model.Error) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(error)
}

// ResponseJSON : to give JSON response
func ResponseJSON(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(data)
}

// ComparePasswords : to compare bcrypt hashed password with given plaintext password.
func ComparePasswords(hashedPassword string, passwordGiven []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), passwordGiven)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

// GenerateToken : When user logs in, we will call GenerateToken() to generate token for login handler using User's credentials.
func GenerateToken(user model.User) (string, error) {
	var err error

	// a jwt
	// header.payload.secret
	secret := os.Getenv("SECRET")

	// create a jwt with claims using User's Email.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"iss":   "course",
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Fatal(err)
	}

	return tokenString, nil
}
