package models

import (
	"boilerplate/database"
	"boilerplate/models/schema"
	"context"
	"errors"
	"log"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

// validateCreditCardData is a function to validate credit card before insert into database
var validateCreditCardData = func(creditCard *schema.CreditCard) (bool, string) {
	// Validate if the credit card has a user id
	if creditCard.UserID == 0 {
		return false, "Credit card need to have a user!"
	}

	// Validate if the credit card has number and more than 15 characters
	if creditCard.Number == "" {
		return false, "Credit card number cannot be empty!"
	} else if len(creditCard.Number) < 16 {
		return false, "Credit card number must be at least 16 characters!"
	}

	// Validate if exist registered credit card with same number
	existCreditCard, _ := schema.CreditCards(schema.CreditCardWhere.Number.EQ(creditCard.Number)).Exists(context.Background(), database.InstanceDB)
	if existCreditCard {
		return false, "There is already a registered credit card with this number, try another number!"
	}

	// Validation passed
	return true, ""
}

// NewCreditCard is a function to insert a single new credit card into database
var NewCreditCard = func(creditCard *schema.CreditCard) (*schema.CreditCard, error) {
	// Validate credit card data to insert
	if valid, messageError := validateCreditCardData(creditCard); !valid {
		return nil, errors.New(messageError)
	}

	// Validate if exist user with creditCard.UserID
	user, _ := schema.FindUser(context.Background(), database.InstanceDB, creditCard.UserID)
	if user == nil {
		return nil, errors.New("not found user")
	}

	// Insert credit card into database
	err := creditCard.Insert(context.Background(), database.InstanceDB, boil.Infer())
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Get new credit card created
	creditCardCreated, err := schema.CreditCards(qm.SQL("select * from users order by id desc")).One(context.Background(), database.InstanceDB)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return creditCardCreated, nil
}

// GetAllCreditCards is a function to return all credit cards registered in database
var GetAllCreditCards = func() ([]*schema.CreditCard, error) {
	allCreditCards, err := schema.CreditCards().All(context.Background(), database.InstanceDB)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return allCreditCards, nil
}

// GetCreditCardByID is a function to return a single credit card
var GetCreditCardByID = func(creditCardId int) (*schema.CreditCard, error) {
	creditCard, err := schema.FindCreditCard(context.Background(), database.InstanceDB, creditCardId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return creditCard, nil
}

// UpdateCreditCard is a function to update data from a single credit card
var UpdateCreditCard = func(creditCardToUpdate *schema.CreditCard) (int64, error) {
	// Validate if exist credit card with creditCardToUpdate.UserID
	creditCard, _ := schema.FindCreditCard(context.Background(), database.InstanceDB, creditCardToUpdate.ID)
	if creditCard == nil {
		return 0, errors.New("not found user")
	}

	// Update credit card with creditCardToUpdate data
	rowsAff, err := creditCardToUpdate.Update(context.Background(), database.InstanceDB, boil.Whitelist("number", "active")) // only update number and active columns
	if err != nil {
		log.Println(err)
		return 0, err
	}

	// Validate if there were lines affected
	if rowsAff < 0 {
		return 0, errors.New("no affected lines")
	}

	// Return affected rows with update
	return rowsAff, nil
}

// DeleteCreditCardByID is a function to delete a single credit card
var DeleteCreditCardByID = func(creditCardID int) (int64, error) {
	// Validate if exist credit card with id equal to creditCardID
	creditCard, _ := schema.FindCreditCard(context.Background(), database.InstanceDB, creditCardID)
	if creditCard == nil {
		return 0, errors.New("not found user")
	}

	// Delete credit card with id equal to creditCardID
	rowsAff, err := creditCard.Delete(context.Background(), database.InstanceDB)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	// Validate if there were lines affected
	if rowsAff < 0 {
		return 0, errors.New("no affected lines")
	}

	// Return affected rows with delete
	return rowsAff, nil
}

// GetAllCreditCardsByUser is a function to return all credit cards by a single user
var GetAllCreditCardsByUser = func(userId int) ([]*schema.CreditCard, error) {
	// Validate if exist user with id equal to userId
	user, _ := schema.FindUser(context.Background(), database.InstanceDB, userId)
	if user == nil {
		return nil, errors.New("not found user")
	}

	// Get all credit cards by user
	creditCards, err := user.CreditCards().All(context.Background(), database.InstanceDB)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return creditCards, nil
}
