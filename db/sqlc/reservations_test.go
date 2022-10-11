package db

import (
	"context"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func createRandomReservation(t *testing.T) {
	arg := CreateReservationParams{
		TenantID:  int32(gofakeit.IntRange(1, 10)),
		RoomID:    int32(gofakeit.IntRange(28, 35)),
		StartDate: time.Now(),
		EndDate:   time.Now().Add(time.Duration(time.Now().Month())),
		Price:     200.00,
		Total:     320.00,
	}

	reservation, err := testQueries.CreateReservation(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, reservation)
	require.Equal(t, arg.Price, reservation.Price)
	require.Equal(t, arg.Total, reservation.Total)
	require.NotZero(t, reservation.StartDate)
	require.NotZero(t, reservation.EndDate)
	require.WithinDuration(t, arg.StartDate, reservation.StartDate, 10*time.Second)
	require.WithinDuration(t, arg.EndDate, reservation.EndDate, 10*time.Second)
}
func TestCreateReservation(t *testing.T) {

	createRandomReservation(t)
}

func createARandomReservation(t *testing.T) Reservation {
	arg := CreateReservationParams{
		TenantID:  int32(gofakeit.IntRange(1, 10)),
		RoomID:    int32(gofakeit.IntRange(28, 35)),
		StartDate: time.Now(),
		EndDate:   time.Now().Add(time.Duration(time.Now().Day())),
		Price:     gofakeit.Float64Range(250, 500),
		Total:     gofakeit.Float64Range(500, 600),
	}

	reservation, err := testQueries.CreateReservation(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, reservation)
	require.Equal(t, arg.Price, reservation.Price)
	require.Equal(t, arg.Total, reservation.Total)
	require.NotZero(t, reservation.StartDate)
	require.NotZero(t, reservation.EndDate)
	require.WithinDuration(t, arg.StartDate, reservation.StartDate, 10*time.Second)
	require.WithinDuration(t, arg.EndDate, reservation.EndDate, 10*time.Second)
	return reservation
}
func TestGetReservations(t *testing.T) {

	reservation1 := createARandomReservation(t)

	reservations, err := testQueries.GetReservations(context.Background(), reservation1.TenantID)

	require.NoError(t, err)

	require.Len(t, reservations, len(reservations))
	for _, returnedRes := range reservations {
		require.NotEmpty(t, returnedRes)
	}

}

//update reservation

func createBRandomReservation(t *testing.T) Reservation {
	arg := CreateReservationParams{
		TenantID:  int32(gofakeit.IntRange(1, 10)),
		RoomID:    int32(gofakeit.IntRange(28, 35)),
		StartDate: time.Now(),
		EndDate:   time.Now().Add(time.Duration(time.Now().Month())),
		Price:     250.00,
		Total:     370.00,
	}

	reservation, err := testQueries.CreateReservation(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, reservation)
	require.Equal(t, arg.Price, reservation.Price)
	require.Equal(t, arg.Total, reservation.Total)
	require.NotZero(t, reservation.StartDate)
	require.NotZero(t, reservation.EndDate)
	require.WithinDuration(t, arg.StartDate, reservation.StartDate, 10*time.Second)
	require.WithinDuration(t, arg.EndDate, reservation.EndDate, 10*time.Second)
	return reservation
}

func TestUpdateReservation(t *testing.T) {

	res1 := createBRandomReservation(t)

	arg := UpdateReservationParams{
		ID:        res1.ID,
		TenantID:  res1.TenantID,
		RoomID:    int32(gofakeit.IntRange(28, 35)),
		StartDate: time.Now().Add(time.Duration(time.Now().Day())),
		EndDate:   time.Now().Add(time.Duration(time.Now().Month())),
		Price:     600.00,
		Total:     650.00,
	}

	updatedRes, err := testQueries.UpdateReservation(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedRes)

	require.Equal(t, arg.ID, updatedRes.ID)
	require.Equal(t, arg.TenantID, updatedRes.TenantID)
	require.Equal(t, arg.RoomID, updatedRes.RoomID)
	require.WithinDuration(t, arg.StartDate, updatedRes.StartDate, 10*time.Second)
	require.WithinDuration(t, arg.EndDate, updatedRes.EndDate, 10*time.Second)
	require.Equal(t, arg.Price, updatedRes.Price)
	require.Equal(t, arg.Total, updatedRes.Total)
}

func TestListReservations(t *testing.T) {

	reservations, err := testQueries.ListReservations(context.Background())

	require.NoError(t, err)
	require.Len(t, reservations, len(reservations))
	for _, returnedRes := range reservations {
		require.NotEmpty(t, returnedRes)
	}

}
