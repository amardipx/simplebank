package api

import (
	"os"
	"testing"
	"time"

	db "github.com/amardipx/simplebank/db/sqlc"
	"github.com/amardipx/simplebank/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		TokenSymmetricKey:   "5XczFbTEhqvdF8d/PfbWbwOJY7d02A8k5GA0Fns03aA=",
		AccessTokenDuration: 15 * time.Minute,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
