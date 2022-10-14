package token

import "time"

// Maker is an interface for managing tokens
type Maker interface {
	// CreateToken creates a new token for a specific email and duration
	CreateToken(email string, duration time.Duration) (string, error)

	// VerifyToken checks the validity of the token
	VerifyToken(token string) (*Payload, error)
}
