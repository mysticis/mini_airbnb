package db

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/mysticis/airbnb_mini/utils"
	"github.com/stretchr/testify/require"
)

func createRandomTenant(t *testing.T) Tenant {
	hashedPassword, err := utils.HashPassword("tenantPass")
	require.NoError(t, err)
	arg := CreateTenantParams{
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		Email:     gofakeit.Email(),
		Phone:     gofakeit.Phone(),
		Password:  hashedPassword,
	}

	tenant, err := testQueries.CreateTenant(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, tenant)

	require.Equal(t, arg.FirstName, tenant.FirstName)
	require.Equal(t, arg.LastName, tenant.LastName)
	require.Equal(t, arg.Email, tenant.Email)
	require.Equal(t, arg.Phone, tenant.Phone)
	err = utils.CheckPassword("tenantPass", hashedPassword)
	require.NoError(t, err)
	return tenant
}
func TestCreateTenant(t *testing.T) {

	createRandomTenant(t)

}
