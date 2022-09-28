package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomImage(t *testing.T) RoomMedium {

	arg := CreateMediaParams{
		RoomID:     4,
		RoomImages: "Test Image",
	}
	roomMedia, err := testQueries.CreateMedia(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, roomMedia)
	require.Equal(t, arg.RoomID, roomMedia.RoomID)
	require.Equal(t, arg.RoomImages, roomMedia.RoomImages)
	return roomMedia
}

func TestCreateRoomImage(t *testing.T) {

	createRandomImage(t)

}

func TestGetRoomImage(t *testing.T) {

	roomImage := createRandomImage(t)

	roomImg, err := testQueries.GetRoomMedia(context.Background(), roomImage.RoomID)
	require.NoError(t, err)
	require.NotEmpty(t, roomImg)
	require.Equal(t, roomImage.RoomID, roomImg.RoomID)
	require.Equal(t, roomImage.RoomImages, roomImg.RoomImages)
}

func TestUpdateRoomMedia(t *testing.T) {

	testRoomMedia := createRandomImage(t)

	arg := UpdateRoomMediaParams{
		ID:         testRoomMedia.ID,
		RoomID:     testRoomMedia.RoomID,
		RoomImages: "Test Image to Update",
	}

	roomMediaUpdated, err := testQueries.UpdateRoomMedia(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, roomMediaUpdated)
	require.Equal(t, testRoomMedia.ID, roomMediaUpdated.ID)
	require.Equal(t, testRoomMedia.RoomID, roomMediaUpdated.RoomID)
	require.Equal(t, arg.RoomImages, roomMediaUpdated.RoomImages)
}

func createRandomImageToDel(t *testing.T) RoomMedium {

	arg := CreateMediaParams{
		RoomID:     5,
		RoomImages: "Test Image to del",
	}
	roomMedia, err := testQueries.CreateMedia(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, roomMedia)
	require.Equal(t, arg.RoomID, roomMedia.RoomID)
	require.Equal(t, arg.RoomImages, roomMedia.RoomImages)
	return roomMedia
}

func TestDeleteRoomMedia(t *testing.T) {

	mediaToDel := createRandomImageToDel(t)

	arg := DeleteRoomMediaParams{
		ID:     mediaToDel.ID,
		RoomID: mediaToDel.RoomID,
	}

	err := testQueries.DeleteRoomMedia(context.Background(), arg)
	require.NoError(t, err)
	media2, err := testQueries.GetRoomMedia(context.Background(), mediaToDel.RoomID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, media2)
}

func TestListRoomMedia(t *testing.T) {

	med1 := createRandomImage(t)

	allMedia, err := testQueries.ListRoomsMedia(context.Background(), med1.RoomID)
	require.NoError(t, err)
	require.Len(t, allMedia, len(allMedia))
	for _, returnedMedia := range allMedia {
		require.NotEmpty(t, returnedMedia)
	}

}
