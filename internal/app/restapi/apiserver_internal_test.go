package restapi

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/danya1off/http-rest-api/internal/app/config"
	"github.com/stretchr/testify/assert"
)

func TestApiServer_HelloHandler(t *testing.T) {
	api := New(config.NewConfig())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	api.helloHandler().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "Hello")
}
