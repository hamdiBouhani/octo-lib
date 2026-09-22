package testutils

import (
	"context"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/hamdiBouhani/octo-lib/httpserver"
	"github.com/hamdiBouhani/octo-lib/logger"
)

type GodogContext struct {
	Server   *gin.Engine
	Recorder *httptest.ResponseRecorder
	Ctx      context.Context
	Token    string
}

func NewGodogContext() *GodogContext {
	gin.SetMode(gin.TestMode)

	// IMPORTANT: real logger to avoid nil panic
	log := logger.New("debug")

	// IMPORTANT: real server builder (middleware-safe)
	srv := httpserver.New(":8080", log, "octo-lib")

	return &GodogContext{
		Server:   srv.Engine(),
		Recorder: httptest.NewRecorder(),
		Ctx:      context.Background(),
	}
}

func (gc *GodogContext) Reset() {
	gc.Recorder = httptest.NewRecorder()
	gc.Ctx = context.Background()
	gc.Token = ""
}
