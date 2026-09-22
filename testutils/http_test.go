package testutils_test

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/hamdiBouhani/octo-lib/testutils"
)

func TestHealthEndpoint(t *testing.T) {
	ts := testutils.NewTestServer()

	ts.Engine.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	rec := ts.PerformRequest("GET", "/health", nil)

	if rec.Code != 200 {
		t.Fatalf("expected 200, got %d", rec.Code)
	}

	var out map[string]string
	testutils.MustDecodeJSON(t, rec, &out)

	if out["status"] != "ok" {
		t.Fatalf("expected ok, got %s", out["status"])
	}
}
