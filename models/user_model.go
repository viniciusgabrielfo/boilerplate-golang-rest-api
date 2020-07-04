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

// validateUserData is a function to validate user before insert into database
var validateUserData = func(user *schema.User) (bool, string) {
	// Validate if the user has a name
	if user.Name == "" {
		return false, "User name cannot be empty!"
	}

	// Validate if the user has email
	if user.Email == "" {
		return false, "User e-mail cannot be empty!"
	}

	// Validate if the user has password and more than 5 characters
	if user.Password == "" {
		return false, "User password cannot be empty!"
	} else if len(user.Password) < 6 {
		return false, "User password must be at least 6 characters!"
	}

	// Validate if exist registered user with same email
	existUser, _ := schema.Users(schema.UserWhere.Email.EQ(user.Email)).Exists(context.Background(), database.InstanceDB)
	if existUser {
		return false, "There is already a registered user with this email, try another email!"
	}

	// Validation passed
	return true, ""
}

// NewUser is a function to insert a single new user into database
var NewUser = func(user *schema.User) (*schema.User, error) {
	// Validate user data to insert
	if valid, messageError := validateUserData(user); !valid {
		return nil, errors.New(messageError)
	}

	// Insert user into database
	err := user.Insert(context.Background(), database.InstanceDB, boil.Infer())
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Get new user created
	userCreated, err := schema.Users(qm.SQL("select id, name, email from users order by id desc")).One(context.Background(), database.InstanceDB)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return userCreated, nil
}

// GetAllUsers is a function to return all users registered in database
var GetAllUsers = func() ([]*schema.User, error) {
	allUsers, err := schema.Users().All(context.Background(), database.InstanceDB)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return allUsers, nil
}

// GetUserByID is a function to return a single user
var GetUserByID = func(userId int) (*schema.User, error) {
	user, err := schema.FindUser(context.Background(), database.InstanceDB, userId, "id", "name", "email") // return only id, name and email columns
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return user, nil
}

// UpdateUser is a function to update data from a single user
var UpdateUser = func(userToUpdate *schema.User) (int64, error) {
	// Validate if exist user with id equal to userId
	user, _ := schema.FindUser(context.Background(), database.InstanceDB, userToUpdate.ID)
	if user == nil {
		return 0, errors.New("not found user")
	}

	// Update user with userToUpdate data
	rowsAff, err := userToUpdate.Update(context.Background(), database.InstanceDB, boil.Whitelist("name", "email")) // only update name and email columns
	if err != nil {
		log.Println(err)
		return 0, err
	}

	if rowsAff < 0 {
		return 0, errors.New("no affected lines")
	}

	// Return affected rows with update
	return rowsAff, nil
}

// DeleteUserByID is a function to delete a single user
var DeleteUserByID = func(userId int) (int64, error) {
	// Validate if exist user with id equal to userId
	user, _ := schema.FindUser(context.Background(), database.InstanceDB, userId)
	if user == nil {
		return 0, errors.New("not found user")
	}

	// Delete user with id equal to userId
	rowsAff, err := user.Delete(context.Background(), database.InstanceDB)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	if rowsAff < 0 {
		return 0, errors.New("no affected lines")
	}

	// Return affected rows with delete
	return rowsAff, nil
}
