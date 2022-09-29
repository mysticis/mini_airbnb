package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/mysticis/airbnb_mini/db/sqlc"
)

type createRoomRequest struct {
	OwnerID              int32     `json:"owner_id" binding:"required"`
	HomeType             []string  `json:"home_type" binding:"required"`
	HomeSize             string    `json:"home_size" binding:"required oneof=15m_squared_or_more 30m_squared_or_more 60m_squared_or_more 90m_squared_or_more"`
	Furnished            bool      `json:"furnished" binding:"required"`
	PrivateBathroom      bool      `json:"private_bathroom" binding:"required"`
	Balcony              bool      `json:"balcony" binding:"required"`
	Garden               bool      `json:"garden" binding:"required"`
	Kitchen              bool      `json:"kitchen" binding:"required"`
	PetsAllowed          bool      `json:"pets_allowed" binding:"required"`
	Parking              bool      `json:"parking" binding:"required"`
	WheelchairAccessible bool      `json:"wheelchair_accessible" binding:"required"`
	Basement             bool      `json:"basement" binding:"required"`
	Amenities            []string  `json:"amenities" binding:"required"`
	SuitableFor          []string  `json:"suitable_for" binding:"required"`
	PublishedAt          time.Time `json:"published_at" binding:"required"`
	Price                float64   `json:"price" binding:"required"`
	CreatedAt            time.Time `json:"created_at" binding:"required"`
	UpdatedAt            time.Time `json:"updated_at" binding:"required"`
	Longitude            float64   `json:"longitude" binding:"required"`
	Latitude             float64   `json:"latitude" binding:"required"`
}

func (server *Server) createRoom(ctx *gin.Context) {

	var req createRoomRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateRoomParams{
		OwnerID:              req.OwnerID,
		HomeType:             req.HomeType,
		HomeSize:             db.HomeSize(req.HomeSize),
		Furnished:            req.Furnished,
		PrivateBathroom:      req.PrivateBathroom,
		Balcony:              req.Balcony,
		Garden:               req.Garden,
		Kitchen:              req.Kitchen,
		PetsAllowed:          req.PetsAllowed,
		Parking:              req.Parking,
		WheelchairAccessible: req.WheelchairAccessible,
		Basement:             req.Basement,
		Amenities:            req.Amenities,
		SuitableFor:          req.SuitableFor,
		Price:                req.Price,
		Longitude:            req.Longitude,
		Latitude:             req.Latitude,
	}

	room, err := server.store.CreateRoom(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, room)
}

//Get a room request

type getRoomByOwnerRequest struct {
	OwnerID int `json:"owner_id,omitempty"`
}

func (server *Server) getRoomByOwner(ctx *gin.Context) {
	var req getRoomByOwnerRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	room, err := server.store.GetRoomByOwner(ctx, int32(req.OwnerID))

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, room)
}

//Update Room

type updateRoomRequest struct {
	OwnerID              int32    `json:"owner_id"`
	HomeType             []string `json:"home_type"`
	HomeSize             string   `json:"home_size" binding:"required oneof=15m_squared_or_more 30m_squared_or_more 60m_squared_or_more 90m_squared_or_more"`
	Furnished            bool     `json:"furnished"`
	PrivateBathroom      bool     `json:"private_bathroom"`
	Balcony              bool     `json:"balcony"`
	Garden               bool     `json:"garden"`
	Kitchen              bool     `json:"kitchen"`
	PetsAllowed          bool     `json:"pets_allowed"`
	Parking              bool     `json:"parking"`
	WheelchairAccessible bool     `json:"wheelchair_accessible"`
	Basement             bool     `json:"basement"`
	Amenities            []string `json:"amenities"`
	SuitableFor          []string `json:"suitable_for"`
	Price                float64  `json:"price"`
	Longitude            float64  `json:"longitude"`
	Latitude             float64  `json:"latitude"`
}

func (server *Server) updateRoomDetails(ctx *gin.Context) {

	var req updateRoomRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateRoomParams{
		HomeType:             req.HomeType,
		HomeSize:             db.HomeSize(req.HomeSize),
		Furnished:            req.Furnished,
		PrivateBathroom:      req.PrivateBathroom,
		Balcony:              req.Balcony,
		Garden:               req.Garden,
		Kitchen:              req.Kitchen,
		PetsAllowed:          req.PetsAllowed,
		Parking:              req.Parking,
		WheelchairAccessible: req.WheelchairAccessible,
		Basement:             req.Basement,
		Amenities:            req.Amenities,
		SuitableFor:          req.SuitableFor,
		Price:                req.Price,
		Longitude:            req.Longitude,
		Latitude:             req.Latitude,
	}

	updatedRoom, err := server.store.UpdateRoom(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, updatedRoom)
}

type listRoomsRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5"`
}

func (server *Server) listRooms(ctx *gin.Context) {

	var req listRoomsRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.ListAllRoomsParams{
		Limit:  int32(req.PageSize),
		Offset: int32(req.PageID-1) * req.PageSize,
	}
	rooms, err := server.store.ListAllRooms(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, rooms)
}
