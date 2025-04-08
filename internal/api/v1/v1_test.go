package v1_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	v2 "gitlab.com/theztd/troll/internal/api/v2"
)

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	v2.RoutesAdd(router.Group("/v1")) // nebo tvůj způsob registrace rout
	return router
}

func TestGetInfo(t *testing.T) {
	router := setupTestRouter()

	req, _ := http.NewRequest(http.MethodGet, "/v2/info", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	// Případně kontrola obsahu JSON:
	// assert.Contains(t, resp.Body.String(), `"version":`)
}

func TestGetStatus(t *testing.T) {
	router := setupTestRouter()

	req, _ := http.NewRequest(http.MethodGet, "/v1/status", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	// Případně kontrola obsahu JSON:
	// assert.Contains(t, resp.Body.String(), `"version":`)
}
