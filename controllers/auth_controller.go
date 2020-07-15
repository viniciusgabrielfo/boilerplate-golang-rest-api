package controllers

import (
	"boilerplate/models"
	u "boilerplate/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// SigninData is a struct that stores auth data
type SigninData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RefreshData is a struct that stores refresh token
type RefreshData struct {
	RefreshToken string `json:"refresh_token"`
}

// Signin is a function to execute login in application and generate JWT "auth token" and "refresh token"
func Signin(w http.ResponseWriter, r *http.Request) {
	signin := &SigninData{}
	err := json.NewDecoder(r.Body).Decode(&signin)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Validates that the email exists and that the password matches
	valid, err := models.Authenticate(signin.Email, signin.Password)
	if err != nil || !valid {
		u.Respond(w, http.StatusBadRequest, u.NewResponse(true, err.Error(), nil))
		return
	}

	// Define token Claims
	tokenClaims := u.Claims{
		Email: signin.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
		},
	}

	// Define refresh token Claims
	refreshTokenClaims := u.Claims{
		Email: signin.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 8).Unix(),
		},
	}

	// Generate token string
	tokenString := u.GenerateNewJwt(tokenClaims, "signin1")
	if tokenString == "" {
		u.Respond(w, http.StatusInternalServerError, u.NewResponse(true, "", nil))
		return
	}

	// Generate refresh token string
	refreshTokenString := u.GenerateNewJwt(refreshTokenClaims, "signin2")
	if refreshTokenString == "" {
		u.Respond(w, http.StatusInternalServerError, u.NewResponse(true, "", nil))
		return
	}

	// Update refresh user token into database
	_, err = models.UpdateRefreshTokenByEmail(signin.Email, refreshTokenString)
	if err != nil {
		u.Respond(w, http.StatusBadRequest, u.NewResponse(true, err.Error(), nil))
		return
	}

	u.Respond(w, http.StatusOK, u.NewResponse(false, "success", map[string]string{
		"token":         tokenString,
		"refresh_token": refreshTokenString,
	}))
}

func Refresh(w http.ResponseWriter, r *http.Request) {
	refreshData := &RefreshData{}

	err := json.NewDecoder(r.Body).Decode(refreshData)
	if err != nil {
		fmt.Println(err)
		u.Respond(w, http.StatusBadRequest, u.NewResponse(true, "malformed json", nil))
		return
	}

	valid, msgError := u.ValidToken(refreshData.RefreshToken, "signin2")
	if !valid || msgError != "" {
		u.Respond(w, http.StatusUnauthorized, u.NewResponse(true, msgError, nil))
		return
	}

	user, err := models.GetUserByToken(refreshData.RefreshToken)
	if err != nil {
		u.Respond(w, http.StatusBadRequest, u.NewResponse(true, err.Error(), nil))
		return
	}

	if user == nil {
		u.Respond(w, http.StatusOK, u.NewResponse(false, "refresh token is invalid", user))
		return
	}

	// Define token Claims
	tokenClaims := u.Claims{
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
		},
	}

	tokenString := u.GenerateNewJwt(tokenClaims, "signin1")

	u.Respond(w, http.StatusOK, u.NewResponse(false, "success", map[string]string{
		"token": tokenString,
	}))
}
