package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/mysticis/airbnb_mini/db/sqlc"
	"github.com/mysticis/airbnb_mini/utils"
)

type createLandlordRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Phone     string `json:"phone" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=6"`
}

type landlordResponse struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

func newLandlordResponse(landlord db.Landlord) userResponse {
	return userResponse{
		FirstName: landlord.FirstName,
		LastName:  landlord.LastName,
		Email:     landlord.Email,
		Phone:     landlord.Phone,
	}
}

func (server *Server) createLandlord(ctx *gin.Context) {

	var req createLandlordRequest

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		log.Fatal("cannot hashpassword: ", err)
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateLandlordParams{
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		Email:          req.Email,
		Phone:          req.Phone,
		HashedPassword: hashedPassword,
	}

	landlord, err := server.store.CreateLandlord(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	res := landlordResponse{
		FirstName: landlord.FirstName,
		LastName:  landlord.LastName,
		Email:     landlord.Email,
		Phone:     landlord.Phone,
	}

	ctx.JSON(http.StatusCreated, res)
}

// list landlords
func (server *Server) listLandlords(ctx *gin.Context) {

	landlords, err := server.store.ListLandlords(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, landlords)
}

//Landlord login

// tenant login
type landlordLoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type landlordLoginResponse struct {
	AccessToken string       `json:"access_token"`
	Landlord    userResponse `json:"landlord"`
}

func (server *Server) loginLandlord(ctx *gin.Context) {
	var req landlordLoginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	landlord, err := server.store.GetLandlord(ctx, req.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = utils.CheckPassword(landlord.HashedPassword, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accessToken, err := server.tokenMaker.CreateToken(req.Email, server.config.AccessTokenDuration)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := landlordLoginResponse{
		AccessToken: accessToken,
		Landlord:    newLandlordResponse(landlord),
	}

	ctx.JSON(http.StatusOK, response)

}
