package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomReservation(t *testing.T) {
	arg := CreateReservationParams{
		TenantID:  2,
		RoomID:    3,
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
		TenantID:  3,
		RoomID:    2,
		StartDate: time.Now(),
		EndDate:   time.Now().Add(time.Duration(time.Now().Day())),
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
	return reservation
}
func TestGetReservation(t *testing.T) {

	reservation1 := createARandomReservation(t)

	reservation2, err := testQueries.GetReservation(context.Background(), reservation1.TenantID)

	require.NoError(t, err)
	require.NotEmpty(t, reservation2)

	require.Equal(t, reservation1.TenantID, reservation2.TenantID)
	require.Equal(t, reservation1.Price, reservation2.Price)
	require.Equal(t, reservation1.Total, reservation2.Total)
	require.WithinDuration(t, reservation1.StartDate, reservation2.StartDate, 25*time.Hour)
	require.WithinDuration(t, reservation1.EndDate, reservation2.EndDate, 30*time.Hour)

}

//update reservation

func createBRandomReservation(t *testing.T) Reservation {
	arg := CreateReservationParams{
		TenantID:  5,
		RoomID:    3,
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
		RoomID:    1,
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
