package token

type Maker interface {
	CreateToken(username, userId string) (string, error)
	VerifyToken(token string) (*Payload, error)
}
