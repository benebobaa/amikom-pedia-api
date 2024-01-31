package token

import "time"

type Maker interface {
	CreateToken(username, userId string, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}
