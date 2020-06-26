package models

import (
	"boilerplate/database"
	"boilerplate/models/schema"
	"context"
	"fmt"

	"github.com/volatiletech/sqlboiler/boil"
)

// NewUser is a function to insert a single new user into database
var NewUser = func(user *schema.User) {
	err := user.Insert(context.Background(), database.InstanceDB, boil.Infer())

	if err != nil {
		fmt.Println(err)
		return
	}
}

var GetAllUsers = func() []*schema.User {
	allUsers, err := schema.Users().All(context.Background(), database.InstanceDB)
	if err != nil {
		fmt.Println(err)
	}

	return allUsers
}

var GetUserById = func(userId int) *schema.User {
	user, _ := schema.FindUser(context.Background(), database.InstanceDB, userId, "id", "name", "email") // return only id, name and email columns
	if user == nil {
		fmt.Println("not found user")
		return user
	}

	return user
}

// UpdateUser is a function to update data from a single user
var UpdateUser = func(userToUpdate *schema.User) {
	user, _ := schema.FindUser(context.Background(), database.InstanceDB, userToUpdate.ID)
	if user == nil {
		fmt.Println("not found user")
		return
	}

	userToUpdate.Update(context.Background(), database.InstanceDB, boil.Whitelist("name", "email")) // only update name and email columns
}

var DeleteUserById = func(userId int) bool {
	user, _ := schema.FindUser(context.Background(), database.InstanceDB, userId)
	if user == nil {
		fmt.Println("not found user")
		return false
	}

	_, err := user.Delete(context.Background(), database.InstanceDB)
	if err != nil {
		fmt.Println("error on delete user")
		return false
	}

	return true
}
