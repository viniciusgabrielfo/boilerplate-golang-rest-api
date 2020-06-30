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
	ID       string `json:"ID"`
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
		return
	}

	userModel := jUser.ConvertToModel()
	models.NewUser(userModel)

	u.Response(w, NewJSONUser(*userModel))
}

var GetAllUsers = func(w http.ResponseWriter, r *http.Request) {
	allUsers := models.GetAllUsers()

	allJSONUsers := make([]JSONUser, 0)

	for _, user := range allUsers {
		allJSONUsers = append(allJSONUsers, NewJSONUser(*user))
	}

	u.Response(w, allJSONUsers)
}

var GetUserByID = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, _ := strconv.Atoi(params["id"])

	u.Response(w, NewJSONUser(*models.GetUserByID(userId)))
}

var UpdateUserByID = func(w http.ResponseWriter, r *http.Request) {
	jUser := &JSONUser{}

	err := json.NewDecoder(r.Body).Decode(jUser)
	if err != nil {
		fmt.Println(err)
		return
	}

	userModel := jUser.ConvertToModel()
	models.UpdateUser(userModel)

	u.Response(w, NewJSONUser(*userModel))
}

var DeleteUserByID = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, _ := strconv.Atoi(params["id"])

	u.Response(w, models.DeleteUserByID(userId))
}

var GetAllCreditCardsByUserID = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, _ := strconv.Atoi(params["id"])

	u.Response(w, models.GetAllCreditCardsByUser(userId))
}
