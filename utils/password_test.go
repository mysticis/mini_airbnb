package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPasswordHash(t *testing.T) {

	password := "secretjruyterttyyyyy"

	hashedPassword, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)
	err = CheckPassword(hashedPassword, password)
	require.NoError(t, err)
	wrongPassword := "wrongPassword"
	err = CheckPassword(wrongPassword, hashedPassword)
	require.Error(t, err)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}
