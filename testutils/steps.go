package testutils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http/httptest"

	"github.com/cucumber/godog"
)

func (gc *GodogContext) iSendRequest(method, path string) error {
	req := httptest.NewRequest(method, path, nil)
	gc.Recorder = httptest.NewRecorder()
	gc.Server.ServeHTTP(gc.Recorder, req)
	return nil
}

func (gc *GodogContext) iSendJSONRequest(method, path string, body *godog.DocString) error {
	req := httptest.NewRequest(method, path, bytes.NewBuffer([]byte(body.Content)))
	req.Header.Set("Content-Type", "application/json")

	gc.Recorder = httptest.NewRecorder()
	gc.Server.ServeHTTP(gc.Recorder, req)
	return nil
}

func (gc *GodogContext) theResponseCodeShouldBe(code int) error {
	if gc.Recorder.Code != code {
		return fmt.Errorf("expected status %d, got %d", code, gc.Recorder.Code)
	}
	return nil
}

func (gc *GodogContext) theJSONShouldContain(expected *godog.DocString) error {
	var actual map[string]interface{}
	if err := json.Unmarshal(gc.Recorder.Body.Bytes(), &actual); err != nil {
		return fmt.Errorf("invalid JSON response: %v", err)
	}

	var exp map[string]interface{}
	if err := json.Unmarshal([]byte(expected.Content), &exp); err != nil {
		return fmt.Errorf("invalid expected JSON: %v", err)
	}

	for k, v := range exp {
		if actual[k] != v {
			return fmt.Errorf("expected JSON key %s=%v, got %v", k, v, actual[k])
		}
	}

	return nil
}

func (gc *GodogContext) RegisterSteps(s *godog.ScenarioContext) {
	s.Step(`^I send a "([^"]*)" request to "([^"]*)"$`, gc.iSendRequest)
	s.Step(`^I send a "([^"]*)" request to "([^"]*)" with JSON:$`, gc.iSendJSONRequest)
	s.Step(`^the response code should be (\d+)$`, gc.theResponseCodeShouldBe)
	s.Step(`^the JSON should contain:$`, gc.theJSONShouldContain)

	// NEW API — BeforeScenario is deprecated
	s.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		gc.Reset()
		return ctx, nil
	})
}
