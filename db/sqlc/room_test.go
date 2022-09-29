package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateRoom(t *testing.T) {

	arg := CreateRoomParams{
		OwnerID:              1,
		HomeType:             []string{"Shared Room", "Apartment"},
		HomeSize:             "30m_squared_or_more",
		Furnished:            true,
		PrivateBathroom:      true,
		Balcony:              true,
		Garden:               true,
		PetsAllowed:          false,
		Parking:              true,
		WheelchairAccessible: false,
		Basement:             false,
		Amenities:            []string{"Dishwasher", "Washing Machine"},
		SuitableFor:          []string{"Students", "Working Professionals"},
		Price:                345.00,
		Longitude:            3.4567,
		Latitude:             6.7896,
	}

	room, err := testQueries.CreateRoom(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, room)

	require.Equal(t, arg.OwnerID, room.OwnerID)
	require.Equal(t, arg.HomeType, room.HomeType)
	require.Equal(t, arg.HomeSize, room.HomeSize)
	require.Equal(t, arg.Furnished, room.Furnished)
	require.Equal(t, arg.PrivateBathroom, room.PrivateBathroom)
	require.Equal(t, arg.Balcony, room.Balcony)
	require.Equal(t, arg.Garden, room.Garden)
	require.Equal(t, arg.PetsAllowed, room.PetsAllowed)
	require.Equal(t, arg.Parking, room.Parking)
	require.Equal(t, arg.WheelchairAccessible, room.WheelchairAccessible)
	require.Equal(t, arg.Basement, room.Basement)
	require.Equal(t, arg.Amenities, room.Amenities)
	require.Equal(t, arg.SuitableFor, room.SuitableFor)
	require.Equal(t, arg.Price, room.Price)
	require.Equal(t, arg.Longitude, room.Longitude)
	require.Equal(t, arg.Latitude, room.Latitude)

}

func createRandomRoom(t *testing.T) Room {
	arg := CreateRoomParams{
		OwnerID:              2,
		HomeType:             []string{"Private Room", "Shared room"},
		HomeSize:             "60m_squared_or_more",
		Furnished:            true,
		PrivateBathroom:      true,
		Balcony:              true,
		Garden:               false,
		PetsAllowed:          false,
		Parking:              true,
		WheelchairAccessible: false,
		Basement:             true,
		Amenities:            []string{"Air Conditioning", "Washing Machine", "Heating"},
		SuitableFor:          []string{"Students", "Working Professionals", "Couples"},
		Price:                250.75,
		Longitude:            8.4567,
		Latitude:             2.7896,
	}

	room, err := testQueries.CreateRoom(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, room)

	require.Equal(t, arg.OwnerID, room.OwnerID)
	require.Equal(t, arg.HomeType, room.HomeType)
	require.Equal(t, arg.HomeSize, room.HomeSize)
	require.Equal(t, arg.Furnished, room.Furnished)
	require.Equal(t, arg.PrivateBathroom, room.PrivateBathroom)
	require.Equal(t, arg.Balcony, room.Balcony)
	require.Equal(t, arg.Garden, room.Garden)
	require.Equal(t, arg.PetsAllowed, room.PetsAllowed)
	require.Equal(t, arg.Parking, room.Parking)
	require.Equal(t, arg.WheelchairAccessible, room.WheelchairAccessible)
	require.Equal(t, arg.Basement, room.Basement)
	require.Equal(t, arg.Amenities, room.Amenities)
	require.Equal(t, arg.SuitableFor, room.SuitableFor)
	require.Equal(t, arg.Price, room.Price)
	require.Equal(t, arg.Longitude, room.Longitude)
	require.Equal(t, arg.Latitude, room.Latitude)
	return room
}
func TestGetRoom(t *testing.T) {

	room1 := createRandomRoom(t)
	room2, err := testQueries.GetRoomByOwner(context.Background(), room1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, room2)

	require.Equal(t, room1.OwnerID, room2.OwnerID)
	require.Equal(t, room1.HomeType, room2.HomeType)
	require.Equal(t, room1.HomeSize, room2.HomeSize)
	require.Equal(t, room1.Furnished, room2.Furnished)
	require.Equal(t, room1.PrivateBathroom, room2.PrivateBathroom)
	require.Equal(t, room1.Balcony, room2.Balcony)
	require.Equal(t, room1.Garden, room2.Garden)
	require.Equal(t, room1.PetsAllowed, room2.PetsAllowed)
	require.Equal(t, room1.Parking, room2.Parking)
	require.Equal(t, room1.WheelchairAccessible, room2.WheelchairAccessible)
	require.Equal(t, room1.Basement, room2.Basement)
	require.Equal(t, room1.Amenities, room2.Amenities)
	require.Equal(t, room1.SuitableFor, room2.SuitableFor)
	require.Equal(t, room1.Price, room2.Price)
	require.Equal(t, room1.Longitude, room2.Longitude)
	require.Equal(t, room1.Latitude, room2.Latitude)

}

func createARandomRoom(t *testing.T) Room {
	arg := CreateRoomParams{
		OwnerID:              3,
		HomeType:             []string{"Private Room", "Shared room"},
		HomeSize:             "90m_squared_or_more",
		Furnished:            true,
		PrivateBathroom:      true,
		Balcony:              true,
		Garden:               false,
		PetsAllowed:          false,
		Parking:              true,
		WheelchairAccessible: false,
		Basement:             true,
		Amenities:            []string{"Air Conditioning", "Washing Machine", "Heating"},
		SuitableFor:          []string{"Students", "Working Professionals", "Couples"},
		Price:                350.75,
		Longitude:            8.4567,
		Latitude:             2.7896,
	}

	room, err := testQueries.CreateRoom(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, room)

	require.Equal(t, arg.OwnerID, room.OwnerID)
	require.Equal(t, arg.HomeType, room.HomeType)
	require.Equal(t, arg.HomeSize, room.HomeSize)
	require.Equal(t, arg.Furnished, room.Furnished)
	require.Equal(t, arg.PrivateBathroom, room.PrivateBathroom)
	require.Equal(t, arg.Balcony, room.Balcony)
	require.Equal(t, arg.Garden, room.Garden)
	require.Equal(t, arg.PetsAllowed, room.PetsAllowed)
	require.Equal(t, arg.Parking, room.Parking)
	require.Equal(t, arg.WheelchairAccessible, room.WheelchairAccessible)
	require.Equal(t, arg.Basement, room.Basement)
	require.Equal(t, arg.Amenities, room.Amenities)
	require.Equal(t, arg.SuitableFor, room.SuitableFor)
	require.Equal(t, arg.Price, room.Price)
	require.Equal(t, arg.Longitude, room.Longitude)
	require.Equal(t, arg.Latitude, room.Latitude)
	return room
}
func TestListRoomsByOwner(t *testing.T) {

	for i := 0; i < 6; i++ {
		createARandomRoom(t)
	}

	arg := ListRoomsByOwnerParams{
		Limit:  3,
		Offset: 3,
	}

	rooms, err := testQueries.ListRoomsByOwner(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, rooms, 3)
	for _, returnedRooms := range rooms {
		require.NotEmpty(t, returnedRooms)
	}

}

func createRandomRoomB(t *testing.T) Room {
	arg := CreateRoomParams{
		OwnerID:              4,
		HomeType:             []string{"Private Room", "Shared room"},
		HomeSize:             "60m_squared_or_more",
		Furnished:            true,
		PrivateBathroom:      true,
		Balcony:              true,
		Garden:               false,
		PetsAllowed:          false,
		Parking:              true,
		WheelchairAccessible: false,
		Basement:             true,
		Amenities:            []string{"Air Conditioning", "Washing Machine", "Heating"},
		SuitableFor:          []string{"Working Professionals", "Couples"},
		Price:                125.75,
		Longitude:            8.4567,
		Latitude:             2.7896,
	}

	room, err := testQueries.CreateRoom(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, room)

	require.Equal(t, arg.OwnerID, room.OwnerID)
	require.Equal(t, arg.HomeType, room.HomeType)
	require.Equal(t, arg.HomeSize, room.HomeSize)
	require.Equal(t, arg.Furnished, room.Furnished)
	require.Equal(t, arg.PrivateBathroom, room.PrivateBathroom)
	require.Equal(t, arg.Balcony, room.Balcony)
	require.Equal(t, arg.Garden, room.Garden)
	require.Equal(t, arg.PetsAllowed, room.PetsAllowed)
	require.Equal(t, arg.Parking, room.Parking)
	require.Equal(t, arg.WheelchairAccessible, room.WheelchairAccessible)
	require.Equal(t, arg.Basement, room.Basement)
	require.Equal(t, arg.Amenities, room.Amenities)
	require.Equal(t, arg.SuitableFor, room.SuitableFor)
	require.Equal(t, arg.Price, room.Price)
	require.Equal(t, arg.Longitude, room.Longitude)
	require.Equal(t, arg.Latitude, room.Latitude)
	return room
}

func TestUpdateRoom(t *testing.T) {

	room := createRandomRoomB(t)

	arg := UpdateRoomParams{
		HomeType:             []string{"Private Room"},
		HomeSize:             "30m_squared_or_more",
		Furnished:            false,
		PrivateBathroom:      true,
		Balcony:              false,
		Garden:               false,
		PetsAllowed:          false,
		Parking:              true,
		WheelchairAccessible: false,
		Basement:             true,
		Amenities:            []string{"Air Conditioning", "Heating"},
		SuitableFor:          []string{"Working Professionals", "Couples"},
		Price:                250.75,
		Longitude:            2.4567,
		Latitude:             6.7896,
	}

	updatedRoom, err := testQueries.UpdateRoom(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedRoom)

	require.Equal(t, room.OwnerID, updatedRoom.OwnerID)
	require.Equal(t, arg.HomeType, updatedRoom.HomeType)
	require.Equal(t, arg.HomeSize, updatedRoom.HomeSize)
	require.Equal(t, arg.Furnished, updatedRoom.Furnished)
	require.Equal(t, arg.PrivateBathroom, updatedRoom.PrivateBathroom)
	require.Equal(t, arg.Balcony, updatedRoom.Balcony)
	require.Equal(t, arg.Garden, updatedRoom.Garden)
	require.Equal(t, arg.PetsAllowed, updatedRoom.PetsAllowed)
	require.Equal(t, arg.Parking, updatedRoom.Parking)
	require.Equal(t, arg.WheelchairAccessible, updatedRoom.WheelchairAccessible)
	require.Equal(t, arg.Basement, updatedRoom.Basement)
	require.Equal(t, arg.Amenities, updatedRoom.Amenities)
	require.Equal(t, arg.SuitableFor, updatedRoom.SuitableFor)
	require.Equal(t, arg.Price, updatedRoom.Price)
	require.Equal(t, arg.Longitude, updatedRoom.Longitude)
	require.Equal(t, arg.Latitude, updatedRoom.Latitude)

}

func createRandomRoomC(t *testing.T) Room {
	arg := CreateRoomParams{
		OwnerID:              4,
		HomeType:             []string{"Private Room", "Shared room"},
		HomeSize:             "60m_squared_or_more",
		Furnished:            true,
		PrivateBathroom:      true,
		Balcony:              true,
		Garden:               false,
		PetsAllowed:          false,
		Parking:              true,
		WheelchairAccessible: false,
		Basement:             true,
		Amenities:            []string{"Air Conditioning", "Washing Machine", "Heating"},
		SuitableFor:          []string{"Working Professionals", "Couples"},
		Price:                125.75,
		Longitude:            8.4567,
		Latitude:             2.7896,
	}

	room, err := testQueries.CreateRoom(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, room)

	require.Equal(t, arg.OwnerID, room.OwnerID)
	require.Equal(t, arg.HomeType, room.HomeType)
	require.Equal(t, arg.HomeSize, room.HomeSize)
	require.Equal(t, arg.Furnished, room.Furnished)
	require.Equal(t, arg.PrivateBathroom, room.PrivateBathroom)
	require.Equal(t, arg.Balcony, room.Balcony)
	require.Equal(t, arg.Garden, room.Garden)
	require.Equal(t, arg.PetsAllowed, room.PetsAllowed)
	require.Equal(t, arg.Parking, room.Parking)
	require.Equal(t, arg.WheelchairAccessible, room.WheelchairAccessible)
	require.Equal(t, arg.Basement, room.Basement)
	require.Equal(t, arg.Amenities, room.Amenities)
	require.Equal(t, arg.SuitableFor, room.SuitableFor)
	require.Equal(t, arg.Price, room.Price)
	require.Equal(t, arg.Longitude, room.Longitude)
	require.Equal(t, arg.Latitude, room.Latitude)
	return room
}

func TestDeleteRoom(t *testing.T) {

	roomToDel := createRandomRoomC(t)
	arg := DeleteRoomParams{
		ID:      roomToDel.ID,
		OwnerID: roomToDel.OwnerID,
	}
	err := testQueries.DeleteRoom(context.Background(), arg)
	require.NoError(t, err)
	room, err := testQueries.GetRoomByOwner(context.Background(), roomToDel.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, room)

}
