package auth

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

var secretKey = []byte("secret-key")

func CreateJWTToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyJWTToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

func AuthenticateRequest(req *http.Request) error {
	tokenString := req.Header.Get("Authorization")

	if tokenString == "" {
		return fmt.Errorf("no token provided")
	}

	tokenString = tokenString[len("Bearer "):]

	err := VerifyJWTToken(tokenString)
	if err != nil {
		return fmt.Errorf("no token provided")
	}

	return nil
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	type User struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var u User

	json.NewDecoder(r.Body).Decode(&u)

	fmt.Printf("DETAILS")
	fmt.Printf(u.Username)
	fmt.Printf(u.Password)

	if u.Username == "lock" && u.Password == "lock" {
		tokenString, err := CreateJWTToken(u.Username)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Errorf("User Not Accepted")
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, tokenString)
		return
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Printf("Invalid Credentials")
	}
}
