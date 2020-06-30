package models

import (
	"boilerplate/database"
	"boilerplate/models/schema"
	"context"
	"fmt"

	"github.com/volatiletech/sqlboiler/boil"
)

// validateCreditCardData is a function to validate credit card before insert into database
var validateCreditCardData = func(creditCard *schema.CreditCard) bool {
	// Validate if the credit card has a user id
	if creditCard.UserID == 0 {
		fmt.Println("Credit card need to have a user!")
		return false
	}

	// Validate if the credit card has number and more than 15 characters
	if creditCard.Number == "" {
		fmt.Println("Credit card number cannot be empty!")
		return false
	} else if len(creditCard.Number) < 16 {
		fmt.Println("Credit card number must be at least 16 characters!")
		return false
	}

	// Validate if exist registered credit card with same number
	existCreditCard, _ := schema.CreditCards(schema.CreditCardWhere.Number.EQ(creditCard.Number)).Exists(context.Background(), database.InstanceDB)
	if existCreditCard {
		fmt.Println("There is already a registered credit card with this number, try another number!")
		return false
	}

	// Validation passed
	return true
}

// NewCreditCard is a function to insert a single new credit card into database
var NewCreditCard = func(creditCard *schema.CreditCard) bool {
	// Validate credit card data to insert
	if !validateCreditCardData(creditCard) {
		return false
	}

	// Validate if exist user with creditCard.UserID
	user, _ := schema.FindUser(context.Background(), database.InstanceDB, creditCard.UserID)
	if user == nil {
		fmt.Println("not found user by id")
		return false
	}

	// Insert credit card into database
	err := creditCard.Insert(context.Background(), database.InstanceDB, boil.Infer())
	if err != nil {
		fmt.Println(err)
	}

	return true
}

// GetAllCreditCards is a function to return all credit cards registered in database
var GetAllCreditCards = func() []*schema.CreditCard {
	allCreditCards, err := schema.CreditCards().All(context.Background(), database.InstanceDB)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return allCreditCards
}

// GetCreditCardByID is a function to return a single credit card
var GetCreditCardByID = func(creditCardId int) *schema.CreditCard {
	creditCard, _ := schema.FindCreditCard(context.Background(), database.InstanceDB, creditCardId)
	if creditCard == nil {
		fmt.Println("not found credit card")
		return nil
	}

	return creditCard
}

// UpdateCreditCard is a function to update data from a single credit card
var UpdateCreditCard = func(creditCardToUpdate *schema.CreditCard) int64 {
	// Validate if exist credit card with creditCardToUpdate.UserID
	creditCard, _ := schema.FindCreditCard(context.Background(), database.InstanceDB, creditCardToUpdate.ID)
	if creditCard == nil {
		fmt.Println("not found credit card")
		return 0
	}

	// Update credit card with creditCardToUpdate data
	rowsAff, err := creditCardToUpdate.Update(context.Background(), database.InstanceDB, boil.Whitelist("number", "active")) // only update number and active columns
	if err != nil {
		fmt.Println(err)
		return 0
	}

	// Return affected rows with update
	return rowsAff
}

// DeleteCreditCardByID is a function to delete a single credit card
var DeleteCreditCardByID = func(creditCardID int) int64 {
	// Validate if exist credit card with id equal to creditCardID
	creditCard, _ := schema.FindCreditCard(context.Background(), database.InstanceDB, creditCardID)
	if creditCard == nil {
		fmt.Println("not found credit card")
		return 0
	}

	// Delete credit card with id equal to creditCardID
	rowsAff, err := creditCard.Delete(context.Background(), database.InstanceDB)
	if err != nil {
		fmt.Println("error on delete credit card")
		return 0
	}

	// Return affected rows with delete
	return rowsAff
}

// GetAllCreditCardsByUser is a function to return all credit cards by a single user
var GetAllCreditCardsByUser = func(userId int) []*schema.CreditCard {
	// Validate if exist user with id equal to userId
	user, _ := schema.FindUser(context.Background(), database.InstanceDB, userId)
	if user == nil {
		fmt.Println("not found user")
		return nil
	}

	// Get all credit cards by user
	creditCards, err := user.CreditCards().All(context.Background(), database.InstanceDB)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return creditCards
}
