package db

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/mysticis/airbnb_mini/utils"
	"github.com/stretchr/testify/require"
)

func createRandomLandord(t *testing.T) Landlord {
	hashedPassword, err := utils.HashPassword("landlordPass")
	require.NoError(t, err)
	arg := CreateLandlordParams{
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		Email:     gofakeit.Email(),
		Phone:     gofakeit.Phone(),
		Password:  hashedPassword,
	}

	landlord, err := testQueries.CreateLandlord(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, landlord)

	require.Equal(t, arg.FirstName, landlord.FirstName)
	require.Equal(t, arg.LastName, landlord.LastName)
	require.Equal(t, arg.Email, landlord.Email)
	require.Equal(t, arg.Phone, landlord.Phone)
	err = utils.CheckPassword("landlordPass", hashedPassword)
	require.NoError(t, err)
	return landlord
}

func TestCreateLandord(t *testing.T) {

	createRandomLandord(t)

}
