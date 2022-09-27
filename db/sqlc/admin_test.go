package db

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func TestCreateAdmin(t *testing.T) {

	arg := CreateAdminParams{
		Name:     gofakeit.Username(),
		Password: gofakeit.Password(true, true, true, true, false, 7),
	}

	admin, err := testQueries.CreateAdmin(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, admin)
	require.Equal(t, arg.Name, admin.Name)
	require.Equal(t, arg.Password, admin.Password)

}
