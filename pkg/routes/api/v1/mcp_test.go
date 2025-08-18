package v1

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestMCPHandler_YAML(t *testing.T) {
	// Use embedded file for testing instead of hardcoded path
	if err := os.Setenv("MCP_FILE", ""); err != nil {
		t.Fatalf("failed to set env: %v", err)
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/mcp", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if err := MCP(c); err != nil {
		t.Fatalf("handler returned error: %v", err)
	}
	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200 got %d", rec.Code)
	}
}

func TestMCPHandler_JSON(t *testing.T) {
	// Use embedded file for testing instead of hardcoded path
	if err := os.Setenv("MCP_FILE", ""); err != nil {
		t.Fatalf("failed to set env: %v", err)
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/mcp", nil)
	req.Header.Set("Accept", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if err := MCP(c); err != nil {
		t.Fatalf("handler returned error: %v", err)
	}
	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200 got %d", rec.Code)
	}
	if got := rec.Header().Get("Content-Type"); got == "" {
		t.Fatalf("missing Content-Type")
	}

	// Ensure body is valid JSON
	var js interface{}
	if err := json.Unmarshal(rec.Body.Bytes(), &js); err != nil {
		t.Fatalf("response body is not valid JSON: %v", err)
	}
}
