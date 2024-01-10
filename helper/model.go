package helper

import (
	"amikom-pedia-api/model/domain"
	"amikom-pedia-api/model/web/register"
	"amikom-pedia-api/model/web/user"
)

func ToUserResponse(userData domain.User) user.ResponseUser {
	return user.ResponseUser{
		UUID:      userData.UUID,
		Username:  userData.Username,
		Email:     userData.Email,
		Nim:       userData.Nim,
		Name:      userData.Name,
		Bio:       userData.Bio,
		CreatedAt: userData.CreatedAt,
		UpdatedAt: userData.UpdatedAt,
	}
}

func ToCategoryResponses(users []domain.User) []user.ResponseUser {
	var userResponses []user.ResponseUser
	for _, category := range users {
		userResponses = append(userResponses, ToUserResponse(category))
	}
	return userResponses
}

func ToRegisterResponse(registerData domain.Register) register.RegisterResponse {
	return register.RegisterResponse{
		ID:        registerData.ID,
		Email:     registerData.Email,
		Nim:       registerData.Nim,
		Name:      registerData.Name,
		CreatedAt: registerData.CreatedAt,
	}
}
