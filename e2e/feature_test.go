package e2e

import (
	"testing"

	"github.com/cucumber/godog"
	"github.com/gin-gonic/gin"
	"github.com/hamdiBouhani/octo-lib/testutils"
)

func TestHealthFeature(t *testing.T) {
	gc := testutils.NewGodogContext()

	// Register your real routes
	gc.Server.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	suite := godog.TestSuite{
		Name: "health",
		ScenarioInitializer: func(sc *godog.ScenarioContext) {
			gc.RegisterSteps(sc)
		},
		Options: &godog.Options{
			Format: "pretty",
			Paths:  []string{"./features/health.feature"},
		},
	}

	if suite.Run() != 0 {
		t.Fatalf("godog tests failed")
	}
}
