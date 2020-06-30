package models

import (
	"boilerplate/database"
	"boilerplate/models/schema"
	"context"
	"fmt"

	"github.com/volatiletech/sqlboiler/boil"
)

// validateUserData is a function to validate user before insert into database
var validateUserData = func(user *schema.User) bool {
	// Validate if the user has a name
	if user.Name == "" {
		fmt.Println("User name cannot be empty!")
		return false
	}

	// Validate if the user has email
	if user.Email == "" {
		fmt.Println("User e-mail cannot be empty!")
		return false
	}

	// Validate if the user has password and more than 5 characters
	if user.Password == "" {
		fmt.Println("User password cannot be empty!")
		return false
	} else if len(user.Password) < 6 {
		fmt.Println("User password must be at least 6 characters!")
		return false
	}

	// Validate if exist registered user with same email
	existUser, _ := schema.Users(schema.UserWhere.Email.EQ(user.Email)).Exists(context.Background(), database.InstanceDB)
	if existUser {
		fmt.Println("There is already a registered user with this email, try another email!")
		return false
	}

	// Validation passed
	return true
}

// NewUser is a function to insert a single new user into database
var NewUser = func(user *schema.User) bool {
	// Validate user data to insert
	if !validateUserData(user) {
		return false
	}

	// Insert user into database
	err := user.Insert(context.Background(), database.InstanceDB, boil.Infer())

	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

// GetAllUsers is a function to return all users registered in database
var GetAllUsers = func() []*schema.User {
	allUsers, err := schema.Users().All(context.Background(), database.InstanceDB)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return allUsers
}

// GetUserByID is a function to retur a single user
var GetUserByID = func(userId int) *schema.User {
	user, _ := schema.FindUser(context.Background(), database.InstanceDB, userId, "id", "name", "email") // return only id, name and email columns
	if user == nil {
		fmt.Println("not found user")
		return nil
	}

	return user
}

// UpdateUser is a function to update data from a single user
var UpdateUser = func(userToUpdate *schema.User) int64 {
	user, _ := schema.FindUser(context.Background(), database.InstanceDB, userToUpdate.ID)
	if user == nil {
		fmt.Println("not found user")
		return 0
	}

	rowsAff, err := userToUpdate.Update(context.Background(), database.InstanceDB, boil.Whitelist("name", "email")) // only update name and email columns
	if err != nil {
		fmt.Println(err)
		return 0
	}

	// Return affected rows with update
	return rowsAff
}

// DeleteUserByID is a function to delete a single user
var DeleteUserByID = func(userId int) int64 {
	user, _ := schema.FindUser(context.Background(), database.InstanceDB, userId)
	if user == nil {
		fmt.Println("not found user")
		return 0
	}

	rowsAff, err := user.Delete(context.Background(), database.InstanceDB)
	if err != nil {
		fmt.Println("error on delete user")
		return 0
	}

	// Return affected rows with delete
	return rowsAff
}
