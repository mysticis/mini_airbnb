package db

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func createRandomLandord(t *testing.T) Landlord {
	arg := CreateLandlordParams{
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		Email:     gofakeit.Email(),
		Phone:     gofakeit.Phone(),
		Password:  gofakeit.Password(true, false, true, true, false, 6),
	}

	landlord, err := testQueries.CreateLandlord(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, landlord)

	require.Equal(t, arg.FirstName, landlord.FirstName)
	require.Equal(t, arg.LastName, landlord.LastName)
	require.Equal(t, arg.Email, landlord.Email)
	require.Equal(t, arg.Phone, landlord.Phone)
	require.Equal(t, arg.Password, landlord.Password)
	return landlord
}

func TestCreateLandord(t *testing.T) {

	createRandomLandord(t)
	createRandomLandord(t)
	createRandomLandord(t)
	createRandomLandord(t)
	createRandomLandord(t)
}
