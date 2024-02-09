package helper

import (
	"amikom-pedia-api/model/domain"
	"amikom-pedia-api/model/web/image"
	"amikom-pedia-api/model/web/login"
	"amikom-pedia-api/model/web/otp"
	"amikom-pedia-api/model/web/post"
	"amikom-pedia-api/model/web/register"
	"amikom-pedia-api/model/web/user"
	"database/sql"
)

func ToUserResponse(userData domain.User, images []domain.Image) user.ResponseUser {
	var imageResponses []image.CreateImageResponse
	for _, image := range images {
		imageResponses = append(imageResponses, ToImageResponse(image))
	}
	return user.ResponseUser{
		UUID:      userData.UUID,
		Username:  userData.Username,
		Email:     userData.Email,
		Nim:       userData.Nim,
		Name:      userData.Name,
		Bio:       userData.Bio,
		Images:    imageResponses,
		CreatedAt: userData.CreatedAt,
		UpdatedAt: userData.UpdatedAt,
	}
}

func ToUserResponses(users []domain.User) []user.ResponseUser {
	var userResponses []user.ResponseUser
	for _, category := range users {
		userResponses = append(userResponses, ToUserResponse(category, []domain.Image{}))
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
		User:        ToUserResponse(user, []domain.Image{}),
	}
}

func ToSetNewPasswordResponse(otpData domain.Otp) user.ForgotPasswordResponse {
	return user.ForgotPasswordResponse{
		RefCode: otpData.RefCode,
	}
}

func ToOtpResponseWithToken(accessToken sql.NullString) otp.CreateResponseWithToken {
	return otp.CreateResponseWithToken{
		AccessToken: accessToken.String,
	}

}

func ToPostResponse(p domain.Post) post.ResponsePost {
	return post.ResponsePost{
		ID:        p.ID,
		Content:   p.Content,
		UserId:    p.UserId,
		RefPostId: p.RefPostId.String,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func ToPostResponses(p []domain.Post, image []domain.Image) []post.ResponsePost {
	var postResponses []post.ResponsePost
	for _, category := range p {
		postResponses = append(postResponses, ToPostResponse(category))
	}
	return postResponses
}

func ToImageResponse(imageData domain.Image) image.CreateImageResponse {
	return image.CreateImageResponse{
		ImageType: imageData.ImageType.String,
		ImageUrl:  imageData.ImageUrl.String,
		FilePath:  imageData.FilePath.String,
		CreatedAt: imageData.CreatedAt,
		UpdatedAt: imageData.UpdatedAt,
	}
}
