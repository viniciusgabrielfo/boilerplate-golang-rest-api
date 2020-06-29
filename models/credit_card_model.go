package models

import (
	"boilerplate/database"
	"boilerplate/models/schema"
	"context"
	"fmt"

	"github.com/volatiletech/sqlboiler/boil"
)

var NewCreditCard = func(creditCard *schema.CreditCard) *schema.CreditCard {

	user, _ := schema.FindUser(context.Background(), database.InstanceDB, creditCard.UserID)
	if user == nil {
		fmt.Println("not found user by id")
		return creditCard
	}

	err := creditCard.Insert(context.Background(), database.InstanceDB, boil.Infer())
	if err != nil {
		fmt.Println(err)
	}

	return creditCard
}

var GetAllCreditCards = func() []*schema.CreditCard {
	allCreditCards, err := schema.CreditCards().All(context.Background(), database.InstanceDB)
	if err != nil {
		fmt.Println(err)
	}

	return allCreditCards
}

var GetCreditCardById = func(creditCardId int) *schema.CreditCard {
	creditCard, _ := schema.FindCreditCard(context.Background(), database.InstanceDB, creditCardId)
	if creditCard == nil {
		fmt.Println("not found credit card")
		return creditCard
	}

	return creditCard
}

var UpdateCreditCard = func(creditCardToUpdate *schema.CreditCard) {
	creditCard, _ := schema.FindCreditCard(context.Background(), database.InstanceDB, creditCardToUpdate.ID)
	if creditCard == nil {
		fmt.Println("not found credit card")
		return
	}

	creditCardToUpdate.Update(context.Background(), database.InstanceDB, boil.Whitelist("number", "active")) // only update number and active columns
}

var DeleteCreditCardByID = func(creditCardId int) bool {
	creditCard, _ := schema.FindCreditCard(context.Background(), database.InstanceDB, creditCardId)
	if creditCard == nil {
		fmt.Println("not found credit card")
		return false
	}

	_, err := creditCard.Delete(context.Background(), database.InstanceDB)
	if err != nil {
		fmt.Println("error on delete credit card")
		return false
	}

	return true
}

var GetAllCreditCardsByUser = func(userId int) []*schema.CreditCard {
	user, _ := schema.FindUser(context.Background(), database.InstanceDB, userId)
	if user == nil {
		fmt.Println("not found user")
	}

	creditCards, _ := user.CreditCards().All(context.Background(), database.InstanceDB)

	return creditCards
}
