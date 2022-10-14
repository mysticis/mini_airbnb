package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/mysticis/airbnb_mini/db/sqlc"
	"github.com/mysticis/airbnb_mini/utils"
	"golang.org/x/crypto/bcrypt"
)

type createTenantRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Phone     string `json:"phone" binding:"required"`
	Password  string `json:"password" binding:"required,min=6"`
}

type userResponse struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

func newUserResponse(tenant db.Tenant) userResponse {
	return userResponse{
		FirstName: tenant.FirstName,
		LastName:  tenant.LastName,
		Email:     tenant.Email,
		Phone:     tenant.Phone,
	}
}

func (server *Server) createTenant(ctx *gin.Context) {

	var req createTenantRequest
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		log.Fatal("cannot hash password:", err)
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateTenantParams{
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		Email:          req.Email,
		Phone:          req.Phone,
		HashedPassword: hashedPassword,
	}

	tenant, err := server.store.CreateTenant(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	response := userResponse{
		FirstName: tenant.FirstName,
		LastName:  tenant.LastName,
		Email:     tenant.Email,
		Phone:     tenant.Phone,
	}

	ctx.JSON(http.StatusCreated, response)
}

// list tenants
func (server *Server) listTenants(ctx *gin.Context) {

	tenants, err := server.store.ListTenants(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, tenants)
}

// tenant login
type tenantLoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type tenantLoginResponse struct {
	AccessToken string       `json:"access_token"`
	Tenant      userResponse `json:"tenant"`
}

func (server *Server) loginTenant(ctx *gin.Context) {
	var req tenantLoginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	tenant, err := server.store.GetTenant(ctx, req.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(tenant.HashedPassword), []byte(req.Password))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accessToken, err := server.tokenMaker.CreateToken(req.Email, server.config.AccessTokenDuration)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := tenantLoginResponse{
		AccessToken: accessToken,
		Tenant:      newUserResponse(tenant),
	}

	ctx.JSON(http.StatusOK, response)

}
