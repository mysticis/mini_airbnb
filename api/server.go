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
	//Rooms
	router.POST("/room/create", server.createRoom)
	router.GET("/room/:owner_id", server.getRoomByOwner)
	router.PUT("/room/update/:id", server.updateRoomDetails)
	router.GET("/rooms", server.listRooms)

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
