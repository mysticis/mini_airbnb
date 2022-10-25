package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/mysticis/airbnb_mini/db/sqlc"
	"github.com/mysticis/airbnb_mini/utils"
)

type createAdminRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Create Admin
func (server *Server) createAdmin(ctx *gin.Context) {

	var req createAdminRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		log.Fatal("cannot hashpassword: ", err)
	}

	arg := db.CreateAdminParams{
		Name:     req.Username,
		Password: hashedPassword,
	}

	admin, err := server.store.CreateAdmin(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, admin)
}

// admin login
type adminLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password" binding:"required,min=6"`
}

type adminLoginResponse struct {
	AccessToken string        `json:"access_token"`
	Admin       adminResponse `json:"admin"`
}

type adminResponse struct {
	Username string `json:"username"`
}

func newAdminResponse(admin db.Admin) adminResponse {
	return adminResponse{
		Username: admin.Name,
	}
}

func (server *Server) adminLogin(ctx *gin.Context) {
	var req adminLoginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	admin, err := server.store.GetAdmin(ctx, req.Username)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = utils.CheckPassword(admin.Password, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accessToken, err := server.adminToken.CreateAdminToken(req.Username, server.config.AccessTokenDuration)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := adminLoginResponse{
		AccessToken: accessToken,
		Admin:       newAdminResponse(admin),
	}

	ctx.JSON(http.StatusOK, response)

}
