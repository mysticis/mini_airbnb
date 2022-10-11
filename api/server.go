package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/mysticis/airbnb_mini/db/sqlc"
)

type Server struct {
	store  *db.SQLStore
	router *gin.Engine
}

func NewServer(store *db.SQLStore) *Server {

	server := &Server{
		store: store,
	}

	router := gin.Default()
	//API routes
	//Landlord
	router.POST("/landlord/create", server.createLandlord)
	router.GET("/landlords", server.listLandlords)
	//Tenant
	router.POST("/tenant/create", server.createTenant)
	router.GET("/tenants", server.listTenants)
	//Rooms
	router.POST("/room/create/:owner_id", server.createRoom) //create a room with an owner identified
	router.GET("/room", server.getRoomByOwner)               //Get a single room with identified owner
	router.PUT("/room/update", server.updateRoomDetails)     //Update a particular room with identified owner
	router.GET("/rooms", server.listRooms)                   //get all available rooms
	router.GET("/roomsbyowner", server.listRoomsByOwner)     //get rooms by owner

	//Reservations
	router.POST("/reservation/create", server.createReservation)
	router.GET("/reservations", server.getReservationsByTenant)
	router.GET("/allreservations", server.listReservations)
	router.DELETE("/reservation", server.deleteReservation)
	router.PUT("/update/reservation", server.updateReservation)

	//Reviews
	router.POST("/reviews/create", server.createReview)  //create a review for a room
	router.GET("/reviews", server.listRoomReviews)       //list all reviews for a room
	router.PUT("/update/review", server.updateReview)    //update a review
	router.DELETE("/delete/review", server.deleteReview) //delete a review
	router.GET("/review", server.getRoomReview)          //get a single review for a room by a user
	server.router = router

	return server
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

// Start runs the http server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
