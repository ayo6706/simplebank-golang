package api

import (
	"os"
	"testing"

	db "github.com/ayo6706/simplebank/db/sqlc"
	"github.com/ayo6706/simplebank/util"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: 15,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}
func TestMain(m *testing.M) {

	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
