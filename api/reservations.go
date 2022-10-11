package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/mysticis/airbnb_mini/db/sqlc"
)

type createReservationRequest struct {
	TenantID  int64     `form:"tenant_id,omitempty"`
	RoomID    int64     `form:"room_id,omitempty"`
	StartDate time.Time `json:"start_date,omitempty"`
	EndDate   time.Time `json:"end_date,omitempty"`
	Price     float64   `json:"price,omitempty"`
	Total     float64   `json:"total,omitempty"`
}

func (server *Server) createReservation(ctx *gin.Context) {
	var req createReservationRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateReservationParams{
		TenantID:  int32(req.TenantID),
		RoomID:    int32(req.RoomID),
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
		Price:     req.Price,
		Total:     req.Total,
	}

	reservation, err := server.store.CreateReservation(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, reservation)
}

// Get reservations by a tenant
type getReservationsRequest struct {
	TenantID int64 `uri:"tenant_id, binding: required"`
}

func (server *Server) getReservationsByTenant(ctx *gin.Context) {
	var req getReservationsRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	reservations, err := server.store.GetReservations(ctx, int32(req.TenantID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, reservations)
}

//List all reservations

func (server *Server) listReservations(ctx *gin.Context) {

	reservations, err := server.store.ListReservations(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, reservations)
}

// delete a reservation by tenant
type deleteReservationRequest struct {
	ID       int64 `form:"id"`
	TenantID int64 `form:"tenant_id"`
}

func (server *Server) deleteReservation(ctx *gin.Context) {

	var req deleteReservationRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.DeleteReservationParams{
		ID:       int32(req.ID),
		TenantID: int32(req.TenantID),
	}
	err := server.store.DeleteReservation(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, "Reservation has been deleted")

}

//update reservation

type updateReservationRequest struct {
	ID        int64     `form:"id,omitempty" json:"id,omitempty"`
	TenantID  int64     `form:"tenant_id,omitempty" json:"tenant_id,omitempty"`
	RoomID    int64     `json:"room_id,omitempty"`
	StartDate time.Time `json:"start_date,omitempty"`
	EndDate   time.Time `json:"end_date,omitempty"`
	Price     float64   `json:"price,omitempty"`
	Total     float64   `json:"total,omitempty"`
}

func (server *Server) updateReservation(ctx *gin.Context) {
	var req updateReservationRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateReservationParams{
		ID:        int32(req.ID),
		TenantID:  int32(req.TenantID),
		RoomID:    int32(req.RoomID),
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
		Price:     req.Price,
		Total:     req.Total,
	}

	updatedRes, err := server.store.UpdateReservation(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, updatedRes)
}
