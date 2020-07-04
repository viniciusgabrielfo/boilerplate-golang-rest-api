package controllers

import (
	"boilerplate/models"
	"boilerplate/models/schema"
	u "boilerplate/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// JSONUser is a struct to receive and response user data on API
type JSONUser struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// ConvertToModel is a function to convert a received JSON from http and convert to model struct type
func (jUser JSONUser) ConvertToModel() *schema.User {
	userIDConverted, _ := strconv.Atoi(jUser.ID)

	user := schema.User{
		ID:       int(userIDConverted),
		Name:     jUser.Name,
		Email:    jUser.Email,
		Password: jUser.Password,
	}

	return &user
}

// NewJSONUser is a function to convert a schema.User type to a JSON structure to return in API
func NewJSONUser(user schema.User) JSONUser {
	return JSONUser{
		ID:       strconv.Itoa(int(user.ID)),
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}

var CreateUser = func(w http.ResponseWriter, r *http.Request) {
	jUser := &JSONUser{}

	err := json.NewDecoder(r.Body).Decode(jUser)
	if err != nil {
		fmt.Println(err)
		u.Respond(w, http.StatusBadRequest, u.NewResponse(true, "malformed json", nil))
		return
	}

	userModel := jUser.ConvertToModel()
	userCreated, err := models.NewUser(userModel)
	if err != nil {
		u.Respond(w, http.StatusBadRequest, u.NewResponse(true, err.Error(), nil))
		return
	}

	u.Respond(w, http.StatusOK, u.NewResponse(false, "success", NewJSONUser(*userCreated)))
}

var GetAllUsers = func(w http.ResponseWriter, r *http.Request) {
	allUsers, err := models.GetAllUsers()
	if err != nil {
		u.Respond(w, http.StatusBadRequest, u.NewResponse(true, err.Error(), nil))
		return
	}

	allJSONUsers := make([]JSONUser, 0)

	for _, user := range allUsers {
		allJSONUsers = append(allJSONUsers, NewJSONUser(*user))
	}

	u.Respond(w, http.StatusOK, u.NewResponse(false, "success", allJSONUsers))
}

var GetUserByID = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, _ := strconv.Atoi(params["id"])

	user, err := models.GetUserByID(userID)
	if err != nil {
		u.Respond(w, http.StatusBadRequest, u.NewResponse(true, err.Error(), nil))
		return
	}

	if user == nil {
		u.Respond(w, http.StatusOK, u.NewResponse(false, "not found user", user))
		return
	}

	u.Respond(w, http.StatusOK, u.NewResponse(false, "success", NewJSONUser(*user)))
}

var UpdateUserByID = func(w http.ResponseWriter, r *http.Request) {
	jUser := &JSONUser{}

	err := json.NewDecoder(r.Body).Decode(jUser)
	if err != nil {
		fmt.Println(err)
		u.Respond(w, http.StatusBadRequest, u.NewResponse(true, "malformed json", nil))
		return
	}

	userModel := jUser.ConvertToModel()
	rowsAff, err := models.UpdateUser(userModel)
	if err != nil {
		u.Respond(w, http.StatusBadRequest, u.NewResponse(true, err.Error(), nil))
		return
	}

	u.Respond(w, http.StatusOK, u.NewResponse(false, "success", rowsAff))
}

var DeleteUserByID = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, _ := strconv.Atoi(params["id"])

	rowsAff, err := models.DeleteUserByID(userID)
	if err != nil {
		u.Respond(w, http.StatusBadRequest, u.NewResponse(true, err.Error(), nil))
		return
	}

	u.Respond(w, http.StatusOK, u.NewResponse(false, "success", rowsAff))
}

var GetAllCreditCardsByUserID = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, _ := strconv.Atoi(params["id"])

	creditCards, err := models.GetAllCreditCardsByUser(userID)
	if err != nil {
		u.Respond(w, http.StatusBadRequest, u.NewResponse(true, err.Error(), nil))
		return
	}

	jCreditCards := make([]JSONCreditCard, 0)

	for _, creditCard := range creditCards {
		jCreditCards = append(jCreditCards, NewJSONCreditCard(*creditCard))
	}

	u.Respond(w, http.StatusOK, u.NewResponse(false, "success", creditCards))
}
