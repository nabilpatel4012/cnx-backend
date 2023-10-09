package token

import (
	"time"
)

type Maker interface {
	//CreateToken creates a new token for a specific email and duration
	CreateToken(email string, duration time.Duration) (string, error)

	//VerifyToken checks if the input token is valid or not
	VerifyToken(token string) (*Payload, error)
}
