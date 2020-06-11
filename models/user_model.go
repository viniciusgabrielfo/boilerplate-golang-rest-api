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

// UpdateUser is a function to update data from a single user
var UpdateUser = func(userToUpdate *schema.User) {
	user, _ := schema.FindUser(context.Background(), database.InstanceDB, userToUpdate.ID)
	if user == nil {
		fmt.Println("not found user")
		return
	}

	userToUpdate.Update(context.Background(), database.InstanceDB, boil.Whitelist("name", "email")) // only update name and email columns
}
