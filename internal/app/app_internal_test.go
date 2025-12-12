package service

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/gorilla/mux"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/require"
// )

// func TestAPIServer_HandleHello(t *testing.T) {
// 	app := New()

// 	router := mux.NewRouter()
// 	app.registerRoutes(router)

// 	rec := httptest.NewRecorder()
// 	req, err := http.NewRequest(http.MethodGet, "/hello", nil)
// 	require.NoError(t, err)

// 	router.ServeHTTP(rec, req)

// 	assert.Equal(t, http.StatusOK, rec.Code)
// 	assert.Equal(t, "Hello", rec.Body.String())
// }
