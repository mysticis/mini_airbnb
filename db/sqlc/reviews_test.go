package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func TestCreateReview(t *testing.T) {

	arg := CreateReviewParams{
		UserID:  3,
		RoomID:  2,
		Comment: gofakeit.Sentence(7),
		Rating:  int32(gofakeit.IntRange(1, 7)),
	}

	review, err := testQueries.CreateReview(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, review)
	require.Equal(t, arg.UserID, review.UserID)
	require.Equal(t, arg.RoomID, review.RoomID)
	require.Equal(t, arg.Comment, review.Comment)
	require.Equal(t, arg.Rating, review.Rating)
}

func createRandomReview(t *testing.T) Review {
	arg := CreateReviewParams{
		UserID:  2,
		RoomID:  3,
		Comment: gofakeit.Sentence(7),
		Rating:  int32(gofakeit.IntRange(1, 5)),
	}

	review, err := testQueries.CreateReview(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, review)
	require.Equal(t, arg.UserID, review.UserID)
	require.Equal(t, arg.RoomID, review.RoomID)
	require.Equal(t, arg.Comment, review.Comment)
	require.Equal(t, arg.Rating, review.Rating)
	return review
}

func TestUpdateReview(t *testing.T) {

	randReview := createRandomReview(t)
	arg := UpdateReviewParams{
		UserID:  randReview.UserID,
		RoomID:  randReview.RoomID,
		Comment: "Test comment",
		Rating:  int32(gofakeit.IntRange(1, 5)),
	}

	review, err := testQueries.UpdateReview(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, review)
	require.Equal(t, randReview.UserID, review.UserID)
	require.Equal(t, arg.UserID, review.UserID)
	require.Equal(t, arg.Comment, review.Comment)
	require.Equal(t, arg.Rating, review.Rating)

}

func TestListReviews(t *testing.T) {

	review := createRandomReview(t)
	reviews, err := testQueries.ListRoomReviews(context.Background(), review.RoomID)
	require.NoError(t, err)
	require.Len(t, reviews, len(reviews))
	for _, allReviews := range reviews {
		require.NotEmpty(t, allReviews)
	}

}

func TestDeleteReview(t *testing.T) {

	revToDel := createRandomReview(t)

	args := GetRoomReviewParams{
		UserID: revToDel.UserID,
		RoomID: revToDel.RoomID,
	}
	arg := DeleteReviewParams{
		RoomID: revToDel.RoomID,
		UserID: revToDel.UserID,
	}
	err := testQueries.DeleteReview(context.Background(), arg)
	require.NoError(t, err)
	review, err := testQueries.GetRoomReview(context.Background(), args)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, review)
}
