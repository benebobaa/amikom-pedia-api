package helper

import (
	"amikom-pedia-api/model/domain"
	"amikom-pedia-api/model/web/login"
	"amikom-pedia-api/model/web/otp"
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

func ToUserResponses(users []domain.User) []user.ResponseUser {
	var userResponses []user.ResponseUser
	for _, category := range users {
		userResponses = append(userResponses, ToUserResponse(category))
	}
	return userResponses
}

func ToRegisterResponse(registerData domain.Register, otpData domain.Otp) register.RegisterResponse {
	return register.RegisterResponse{
		ID:        registerData.ID,
		Email:     registerData.Email,
		Nim:       registerData.Nim,
		Name:      registerData.Name,
		CreatedAt: registerData.CreatedAt,
		RefCode:   otpData.RefCode,
	}
}

func ToOtpResponse(otpData domain.Otp) otp.CreateResponseOTP {
	return otp.CreateResponseOTP{
		RefCode: otpData.RefCode,
	}
}

func ToLoginResponse(user domain.User, accessToken string) login.LoginResponse {
	return login.LoginResponse{
		AccessToken: accessToken,
		User:        ToUserResponse(user),
	}
}

func ToSetNewPasswordResponse(otpData domain.Otp) user.ForgotPasswordResponse {
	return user.ForgotPasswordResponse{
		RefCode: otpData.RefCode,
	}
}
