package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var ErrExpiredToken = errors.New("token has expired")

type Payload struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Email     string    `json:"email,omitempty"`
	IssuedAt  time.Time `json:"issued_at,omitempty"`
	ExpiredAt time.Time `json:"expired_at,omitempty"`
}

// NewPayload creates a new token payload with a specific email and duration
func NewPayload(email string, duration time.Duration) (*Payload, error) {

	uuid, err := uuid.NewRandom()

	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        uuid,
		Email:     email,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	return payload, nil
}

//Valid checks if the token payload is valid or not

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
