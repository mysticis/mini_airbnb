package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var ErrExpiredAdminToken = errors.New("admin token has expired")

type AdminPayload struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Username  string    `json:"username,omitempty"`
	IssuedAt  time.Time `json:"issued_at,omitempty"`
	ExpiredAt time.Time `json:"expired_at,omitempty"`
}

// NewPayload creates a new token payload with a specific email and duration
func NewAdminPayload(username string, duration time.Duration) (*AdminPayload, error) {

	uuid, err := uuid.NewRandom()

	if err != nil {
		return nil, err
	}

	adminPayload := &AdminPayload{
		ID:        uuid,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	return adminPayload, nil
}

//Valid checks if the token payload is valid or not

func (adminPayload *AdminPayload) Valid() error {
	if time.Now().After(adminPayload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
