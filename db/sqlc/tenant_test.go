package db

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func createRandomTenant(t *testing.T) Tenant {
	arg := CreateTenantParams{
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		Email:     gofakeit.Email(),
		Phone:     gofakeit.Phone(),
		Password:  gofakeit.Password(true, true, true, true, false, 6),
	}

	tenant, err := testQueries.CreateTenant(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, tenant)

	require.Equal(t, arg.FirstName, tenant.FirstName)
	require.Equal(t, arg.LastName, tenant.LastName)
	require.Equal(t, arg.Email, tenant.Email)
	require.Equal(t, arg.Phone, tenant.Phone)
	return tenant
}
func TestCreateTenant(t *testing.T) {

	createRandomTenant(t)
	createRandomTenant(t)
	createRandomTenant(t)
	createRandomTenant(t)
	createRandomTenant(t)
}
