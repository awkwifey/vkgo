package API

import (
	"fmt"
)

func(config *Options) GetUser(userID int, fields... string) *ResponseUsers {
	var (
		body	= config.HeaderExecute("users.get", fmt.Sprintf("user_id=%d&fields=sex", userID))
		users	= ResponseUsers{}
	)
	users.Bot	= config
	users.Body	= body

	return &users
}
/*
func(config *Options) GetUsers(usersID []int) ResponseUsers {
	var ids string
	for _, id := range usersID {
		ids += fmt.Sprintf("%d", id) + "%2C+"
	}

	body := config.HeaderExecute("users.get", fmt.Sprintf("user_ids=%s", ids))

	var users response.Users
	json.Unmarshal(body, &users)

	if users.Error.Code != 0 {
		message := users.Error.Message
		return []response.User{}, fmt.Errorf(message)
	} else if len(users.Users) < 1 {
		return []response.User{}, nil
	} else {
		return users.Users, nil
	}
}

func(config *Options) GetUserGender(userID int) (int, error) {
	user, err := config.GetUser(userID)
	if err != nil {
		return 0, err
	}
	return user.Gender, nil
}

func(config *Options) GetUserName(userID int) (string, error) {
	user, err := config.GetUser(userID)
	if err != nil {
		return "", err
	}
	return user.Name, nil
}

func(config *Options) GetUsersName(usersID []int) ([]string, error) {
	users, err := config.GetUsers(usersID)
	if err != nil {
		return []string{""}, err
	}

	var names []string
	for _, user := range users {
		names = append(names, user.Name)
	}

	return names, nil
}

func(config *Options) GetUserSurname(userID int) (string, error) {
	user, err := config.GetUser(userID)
	if err != nil {
		return "", err
	}
	return user.Surname, nil
}

func(config *Options) GetUsersSurname(usersID []int) ([]string, error) {
	users, err := config.GetUsers(usersID)
	if err != nil {
		return []string{""}, err
	}

	var surnames []string
	for _, user := range users {
		surnames = append(surnames, user.Surname)
	}

	return surnames, nil
}

var nameCases = map[string]string{
	"nominative":	"nom",
	"genitive":		"gen",
	"dative":		"dat",
	"accusative":	"acc",
	"instrumental":	"ins",
}*/