package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
)

var ErrInvalidAdminToken = errors.New("invalid admin token")

type PasetoAdminMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

//NewPasetoMaker creates a new PasetoMaker

func NewPasetoAdminMaker(symmetricKey string) (AdminMaker, error) {

	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid keysize: must be exactly %d characters: ", chacha20poly1305.KeySize)
	}
	adminMaker := &PasetoAdminMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}
	return adminMaker, nil
}

// CreateToken creates a new token for a specific email and duration
func (maker *PasetoAdminMaker) CreateAdminToken(name string, duration time.Duration) (string, error) {
	adminPayload, err := NewAdminPayload(name, duration)

	if err != nil {
		return "", err
	}

	return maker.paseto.Encrypt(maker.symmetricKey, adminPayload, nil)
}

// VerifyToken checks the validity of the token
func (maker *PasetoAdminMaker) VerifyAdminToken(token string) (*AdminPayload, error) {

	adminPayload := &AdminPayload{}

	err := maker.paseto.Decrypt(token, maker.symmetricKey, adminPayload, nil)

	if err != nil {
		return nil, ErrInvalidAdminToken
	}

	err = adminPayload.Valid()

	if err != nil {
		return nil, err
	}

	return adminPayload, nil
}
