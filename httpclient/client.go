package httpclient

import (
	"context"
	"net/http"
)

type Client struct {
	http.Client
}

func New() *Client {
	return &Client{http.Client{}}
}

func (c *Client) DoWithRequestID(ctx context.Context, req *http.Request) (*http.Response, error) {
	reqID := ctx.Value("request_id")
	if reqID != nil {
		req.Header.Set(
			"X-Request-ID", reqID.(string),
		)
	}
	return c.Do(req)
}
