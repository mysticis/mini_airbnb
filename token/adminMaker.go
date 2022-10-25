package token

import "time"

// Maker is an interface for managing tokens
type AdminMaker interface {
	// CreateToken creates a new token for a specific email and duration
	CreateAdminToken(username string, duration time.Duration) (string, error)

	// VerifyToken checks the validity of the token
	VerifyAdminToken(token string) (*AdminPayload, error)
}
