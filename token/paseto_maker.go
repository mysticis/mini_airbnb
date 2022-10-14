package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
)

var ErrInvalidToken = errors.New("invalid token")

type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

//NewPasetoMaker creates a new PasetoMaker

func NewPasetoMaker(symmetricKey string) (Maker, error) {

	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid keysize: must be exactly %d characters: ", chacha20poly1305.KeySize)
	}
	maker := &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}
	return maker, nil
}

// CreateToken creates a new token for a specific email and duration
func (maker *PasetoMaker) CreateToken(email string, duration time.Duration) (string, error) {
	payload, err := NewPayload(email, duration)

	if err != nil {
		return "", err
	}

	return maker.paseto.Encrypt(maker.symmetricKey, payload, nil)
}

// VerifyToken checks the validity of the token
func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {

	payload := &Payload{}

	err := maker.paseto.Decrypt(token, maker.symmetricKey, payload, nil)

	if err != nil {
		return nil, ErrInvalidToken
	}

	err = payload.Valid()

	if err != nil {
		return nil, err
	}

	return payload, nil
}
