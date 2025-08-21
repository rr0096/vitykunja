package v1

import (
	_ "embed"
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"sync"

	"code.vikunja.io/api/pkg/log"
	"github.com/labstack/echo/v4"
	yamlv3 "gopkg.in/yaml.v3"
)

//go:embed mcp_embedded.yaml
var embeddedMCP []byte

var (
	mcpOnce sync.Once
	mcpData []byte
	mcpErr  error
)

func loadMCP() {
	// Prefer explicit override
	if p := os.Getenv("MCP_FILE"); p != "" {
		data, err := os.ReadFile(p)
		if err != nil {
			mcpErr = err
			return
		}
		mcpData = data
		return
	}

	// fallback to embedded
	if len(embeddedMCP) > 0 {
		mcpData = embeddedMCP
		return
	}

	mcpErr = os.ErrNotExist
}

// MCP serves the MCP specification YAML (docs/mcp.yaml).
func MCP(c echo.Context) error {
	mcpOnce.Do(loadMCP)
	if mcpErr != nil {
		log.Error(mcpErr.Error())
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(mcpErr)
	}
	// Support content negotiation: if Accept contains application/json, convert YAML -> JSON
	accept := c.Request().Header.Get("Accept")
	if strings.Contains(accept, "application/json") {
		var v interface{}
		if err := yamlv3.Unmarshal(mcpData, &v); err != nil {
			log.Error(err.Error())
			return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
		}
		js, err := json.MarshalIndent(v, "", "  ")
		if err != nil {
			log.Error(err.Error())
			return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
		}
		return c.Blob(http.StatusOK, "application/json; charset=utf-8", js)
	}

	return c.Blob(http.StatusOK, "text/yaml; charset=utf-8", mcpData)
}
