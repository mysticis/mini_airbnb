package api

import (
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/mysticis/airbnb_mini/db/sqlc"
	"github.com/mysticis/airbnb_mini/utils"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store) *Server {

	config := utils.Config{
		TokenSymmetricKey:   "abcdefghijklmnopqrstuvwxyznfgrte",
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)
	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
