package login

import "amikom-pedia-api/model/web/user"

type LoginResponse struct {
	AccessToken string            `json:"access_token"`
	User        user.ResponseUser `json:"user"`
}
