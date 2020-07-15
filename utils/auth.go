package utils

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// Claims is a struct to represent Claims for JWF
type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// IsAuthorizedMiddleware is a middleware function to validate JWF token from http request
func IsAuthorizedMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authToken := r.Header.Get("Authorization")
		if authToken == "" {
			Respond(w, http.StatusUnauthorized, NewResponse(true, "missing auth token", nil))
			return
		}

		extractedToken := strings.Split(authToken, "Bearer ")
		if len(extractedToken) != 2 {
			Respond(w, http.StatusBadRequest, NewResponse(true, "auth token is malformed", nil))
			return
		}

		valid, msgError := ValidToken(extractedToken[1], "signin1")
		if !valid || msgError != "" {
			Respond(w, http.StatusUnauthorized, NewResponse(true, msgError, nil))
			return
		}

		next.ServeHTTP(w, r)
	})
}

// GenerateNewJwt is a function to generate new JWF token in string format
func GenerateNewJwt(claims Claims, typeKey string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// There will be two types of JWF key, "signin1" for auth token and "signin2" for refres token
	token.Header["type_key"] = typeKey

	var signKey []byte = nil
	if typeKey == "signin1" {
		signKey = []byte(os.Getenv("jwtKey"))
	} else {
		signKey = []byte(os.Getenv("jwtRefresKey"))
	}

	tokenString, err := token.SignedString(signKey)
	if err != nil {
		log.Println(err)
		return ""
	}

	return tokenString
}

// ValidToken is a function to validate auth or refresh token
func ValidToken(tokenString string, typeKey string) (bool, string) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Valid if is auth or refresh token
		if typeKey == "signin1" {
			return []byte(os.Getenv("jwtKey")), nil
		}
		return []byte(os.Getenv("jwtRefresKey")), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return false, "token is malformed"
		}

		return false, "token is invalid"
	}

	if !token.Valid {
		return false, "token is invalid"
	}

	return true, ""
}
