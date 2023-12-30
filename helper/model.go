package helper

import (
	"amikom-pedia-api/model/domain"
	"amikom-pedia-api/model/web/user"
)

func ToUserResponse(userData domain.User) user.ResponseUser {
	return user.ResponseUser{
		Id:          userData.Id,
		Username:    userData.Username,
		DisplayName: userData.DisplayName,
		Email:       userData.Email,
		Password:    userData.Password,
		CreatedAt:   userData.CreatedAt,
		UpdateAt:    userData.UpdateAt,
	}
}

func ToCategoryResponses(users []domain.User) []user.ResponseUser {
	var userResponses []user.ResponseUser
	for _, category := range users {
		userResponses = append(userResponses, ToUserResponse(category))
	}
	return userResponses
}
