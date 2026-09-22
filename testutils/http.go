package testutils

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

// TestServer wraps Gin + recorder + helpers.
type TestServer struct {
	Engine   *gin.Engine
	Recorder *httptest.ResponseRecorder
}

// NewTestServer creates a Gin test server with a fresh recorder.
func NewTestServer() *TestServer {
	gin.SetMode(gin.TestMode)

	r := gin.New()
	w := httptest.NewRecorder()

	return &TestServer{
		Engine:   r,
		Recorder: w,
	}
}

// PerformRequest executes an HTTP request against the Gin engine.
func (ts *TestServer) PerformRequest(method, path string, body interface{}) *httptest.ResponseRecorder {
	var req *http.Request

	if body != nil {
		jsonBytes, _ := json.Marshal(body)
		req = httptest.NewRequest(method, path, bytes.NewBuffer(jsonBytes))
		req.Header.Set("Content-Type", "application/json")
	} else {
		req = httptest.NewRequest(method, path, nil)
	}

	ts.Recorder = httptest.NewRecorder()
	ts.Engine.ServeHTTP(ts.Recorder, req)

	return ts.Recorder
}

// DecodeJSON decodes the JSON response body into a struct.
func DecodeJSON[T any](rec *httptest.ResponseRecorder, out *T) error {
	return json.Unmarshal(rec.Body.Bytes(), out)
}

// MustDecodeJSON fails the test immediately if JSON cannot be decoded.
func MustDecodeJSON[T any](t TestingT, rec *httptest.ResponseRecorder, out *T) {
	t.Helper()
	if err := DecodeJSON(rec, out); err != nil {
		t.Fatalf("failed to decode JSON: %v", err)
	}
}

// TestingT is the minimal interface needed from *testing.T.
type TestingT interface {
	Helper()
	Fatalf(format string, args ...interface{})
}

// NewJSONRequest creates a JSON request for manual testing.
func NewJSONRequest(method, path string, body interface{}) *http.Request {
	jsonBytes, _ := json.Marshal(body)
	req := httptest.NewRequest(method, path, bytes.NewBuffer(jsonBytes))
	req.Header.Set("Content-Type", "application/json")
	return req
}

// NewEmptyRequest creates a request without a body.
func NewEmptyRequest(method, path string) *http.Request {
	return httptest.NewRequest(method, path, nil)
}
