package api

import (
	"log"

	"github.com/gin-gonic/gin"
	db "github.com/mysticis/airbnb_mini/db/sqlc"
	"github.com/mysticis/airbnb_mini/token"
	"github.com/mysticis/airbnb_mini/utils"
)

type Server struct {
	config     utils.Config
	store      db.Store
	router     *gin.Engine
	tokenMaker token.Maker
	adminToken token.AdminMaker
}

func NewServer(config utils.Config, store db.Store) (*Server, error) {

	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		log.Fatal("cannot create token maker:", err)
	}

	adminTokenMaker, err := token.NewPasetoAdminMaker(config.TokenSymmetricKeyAdmin)
	if err != nil {
		log.Fatal("cannot create token maker for admin:", err)
	}
	server := &Server{
		store:      store,
		tokenMaker: tokenMaker,
		config:     config,
		adminToken: adminTokenMaker,
	}
	server.setupRouter()
	return server, nil
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

// Start runs the http server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func (server *Server) setupRouter() {

	router := gin.Default()
	//API routes
	//Create and Login Routes for Landlords and Tenants
	//Admin Login
	router.POST("/admin/create", server.createAdmin)
	router.POST("/admin/login", server.adminLogin)

	//Landlord APIs
	router.POST("/landlord/create", server.createLandlord)
	router.POST("/tenant/create", server.createTenant)
	//Tenant APIs
	router.POST("/tenant/login", server.loginTenant)
	router.POST("/landlord/login", server.loginLandlord)

	//for admin use
	adminAuthRoutes := router.Group("/").Use(adminAuthMiddleware(server.adminToken))
	//For Landlords and tenants
	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	adminAuthRoutes.GET("/tenants", server.listTenants)

	adminAuthRoutes.GET("/landlords", server.listLandlords) //list landlords

	//Rooms
	authRoutes.POST("/room/create/:owner_id", server.createRoom) //create a room with an owner identified
	authRoutes.GET("/room", server.getRoomByOwner)               //Get a single room with identified owner
	authRoutes.PUT("/room/update", server.updateRoomDetails)     //Update a particular room with identified owner
	adminAuthRoutes.GET("/rooms", server.listRooms)              //get all available rooms
	authRoutes.GET("/roomsbyowner", server.listRoomsByOwner)     //get rooms by owner

	//Reservations
	authRoutes.POST("/reservation/create", server.createReservation) //create a reservation
	authRoutes.GET("/reservations", server.getReservationsByTenant)  //get reservations for a tenant
	adminAuthRoutes.GET("/allreservations", server.listReservations) //list all reservations
	authRoutes.DELETE("/reservation", server.deleteReservation)      //delete a reservation
	authRoutes.PUT("/update/reservation", server.updateReservation)  //update a reservation

	//Reviews
	authRoutes.POST("/reviews/create", server.createReview)  //create a review for a room
	authRoutes.GET("/reviews", server.listRoomReviews)       //list all reviews for a room
	authRoutes.PUT("/update/review", server.updateReview)    //update a review
	authRoutes.DELETE("/delete/review", server.deleteReview) //delete a review
	authRoutes.GET("/review", server.getRoomReview)          //get a single review for a room by a user
	server.router = router

}
