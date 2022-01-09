package services

import (
	"github.com/Lakshay05/go_users_api/domain/users"
	"github.com/Lakshay05/go_users_api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	//Validate
	if err := user.Validate(); err != nil {
		return nil, err
	}
	//Save to database
	if err := user.Save(); err != nil {
		return nil, err

	}

	return &user, nil
}

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func UpdateUser(user users.User, isPartial bool) (*users.User, *errors.RestErr) {
	current, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.FirstName
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email

	}

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func DeleteUser(userId int64) *errors.RestErr {
	user := &users.User{Id: userId}
	return user.Delete()
}
