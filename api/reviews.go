package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/mysticis/airbnb_mini/db/sqlc"
)

type createReviewRequest struct {
	TenantID int32  `form:"tenant_id,omitempty"`
	RoomID   int32  `form:"room_id,omitempty"`
	Comment  string `json:"comment,omitempty"`
	Rating   int32  `json:"rating,omitempty"`
}

func (server *Server) createReview(ctx *gin.Context) {

	var req createReviewRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateReviewParams{
		UserID:  req.TenantID,
		RoomID:  req.RoomID,
		Comment: req.Comment,
		Rating:  req.Rating,
	}

	review, err := server.store.CreateReview(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, review)
}

//list room reviews

type listRoomReviewsRequest struct {
	RoomID int32 `uri:"room_id,min=1"`
}

func (server *Server) listRoomReviews(ctx *gin.Context) {

	var req listRoomReviewsRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	revs, err := server.store.ListRoomReviews(ctx, req.RoomID)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, revs)

}

// update a review
type updateReviewRequest struct {
	UserID  int32  `form:"user_id,binding:required"`
	RoomID  int32  `form:"room_id,binding:required"`
	Comment string `json:"comment,omitempty"`
	Rating  int32  `json:"rating,omitempty"`
}

func (server *Server) updateReview(ctx *gin.Context) {
	var req updateReviewRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateReviewParams{
		UserID:  req.UserID,
		RoomID:  req.RoomID,
		Comment: req.Comment,
		Rating:  req.Rating,
	}
	updatedReview, err := server.store.UpdateReview(ctx, arg)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, updatedReview)
}

// Delete reviews
type deleteReviewsRequest struct {
	UserID int32 `form:"user_id,omitempty"`
	RoomID int32 `form:"room_id,omitempty"`
}

func (server *Server) deleteReview(ctx *gin.Context) {

	var req deleteReviewsRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.DeleteReviewParams{
		RoomID: req.RoomID,
		UserID: req.UserID,
	}

	err := server.store.DeleteReview(ctx, arg)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, "Review has been deleted")
}

//get a single room review

type getRoomReviewRequest struct {
	UserID int32 `form:"user_id"`
	RoomID int32 `form:"room_id"`
}

func (server *Server) getRoomReview(ctx *gin.Context) {
	var req getRoomReviewRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetRoomReviewParams{
		UserID: req.UserID,
		RoomID: req.RoomID,
	}

	rev, err := server.store.GetRoomReview(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, rev)
}
