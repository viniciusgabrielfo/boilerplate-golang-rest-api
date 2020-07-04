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

// JSONCreditCard is a struct to receive and response credit card data on API
type JSONCreditCard struct {
	ID     string `json:"id"`
	Number string `json:"number"`
	Active bool   `json:"active"`
	UserID string `json:"user_id"`
}

// ConvertToModel is a function to convert a received JSON from http and convert to model struct type
func (jCreditCard JSONCreditCard) ConvertToModel() *schema.CreditCard {
	creditCardConvertedID, _ := strconv.Atoi(jCreditCard.ID)
	userConvertedID, _ := strconv.Atoi(jCreditCard.UserID)

	creditCard := schema.CreditCard{
		ID:     creditCardConvertedID,
		Number: jCreditCard.Number,
		Active: jCreditCard.Active,
		UserID: userConvertedID,
	}

	return &creditCard
}

// NewJSONCreditCard is a function to convert a schema.CreditCard to a JSON strucutre to return in API
func NewJSONCreditCard(creditCard schema.CreditCard) JSONCreditCard {
	return JSONCreditCard{
		ID:     strconv.Itoa(int(creditCard.ID)),
		Number: creditCard.Number,
		Active: creditCard.Active,
		UserID: strconv.Itoa(int(creditCard.UserID)),
	}
}

var CreateCreditCard = func(w http.ResponseWriter, r *http.Request) {
	jCreditCard := &JSONCreditCard{}

	err := json.NewDecoder(r.Body).Decode(jCreditCard)
	if err != nil {
		fmt.Println(err)
		u.Respond(w, http.StatusBadRequest, u.NewResponse(true, "malformed json", nil))
		return
	}

	creditCardModel := jCreditCard.ConvertToModel()
	creditCardCreated, err := models.NewCreditCard(creditCardModel)
	if err != nil {
		u.Respond(w, http.StatusBadRequest, u.NewResponse(true, err.Error(), nil))
		return
	}

	u.Respond(w, http.StatusCreated, u.NewResponse(false, "success", NewJSONCreditCard(*creditCardCreated)))
}

var GetAllCreditCards = func(w http.ResponseWriter, r *http.Request) {
	allCreditCards, err := models.GetAllCreditCards()
	if err != nil {
		u.Respond(w, http.StatusBadRequest, u.NewResponse(true, err.Error(), nil))
		return
	}

	allJSONCreditCards := make([]JSONCreditCard, 0)

	for _, creditCard := range allCreditCards {
		allJSONCreditCards = append(allJSONCreditCards, NewJSONCreditCard(*creditCard))
	}

	u.Respond(w, http.StatusOK, u.NewResponse(false, "success", allJSONCreditCards))
}

var GetCreditCardByID = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	creditCardByID, _ := strconv.Atoi(params["id"])

	creditCard, err := models.GetCreditCardByID(creditCardByID)
	if err != nil {
		u.Respond(w, http.StatusBadRequest, u.NewResponse(true, err.Error(), nil))
		return
	}

	if creditCard == nil {
		u.Respond(w, http.StatusOK, u.NewResponse(false, "not found credit card", creditCard))
		return
	}

	u.Respond(w, http.StatusOK, u.NewResponse(false, "success", NewJSONCreditCard(*creditCard)))
}

var UpdateCreditCardByID = func(w http.ResponseWriter, r *http.Request) {
	jCreditCard := &JSONCreditCard{}

	err := json.NewDecoder(r.Body).Decode(jCreditCard)
	if err != nil {
		fmt.Println(err)
		u.Respond(w, http.StatusBadRequest, u.NewResponse(true, "malformed json", nil))
		return
	}

	creditCardModel := jCreditCard.ConvertToModel()
	rowsAff, err := models.UpdateCreditCard(creditCardModel)
	if err != nil {
		u.Respond(w, http.StatusBadRequest, u.NewResponse(true, err.Error(), nil))
		return
	}

	u.Respond(w, http.StatusOK, u.NewResponse(false, "success", rowsAff))
}

var DeleteCreditCardByID = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	creditCardID, _ := strconv.Atoi(params["id"])

	rowsAff, err := models.DeleteCreditCardByID(creditCardID)
	if err != nil {
		u.Respond(w, http.StatusBadRequest, u.NewResponse(true, err.Error(), nil))
	}

	u.Respond(w, http.StatusOK, u.NewResponse(false, "success", rowsAff))
}
