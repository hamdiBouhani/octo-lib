Feature: Health endpoint

  Scenario: Check health
    When I send a "GET" request to "/health"
    Then the response code should be 200
    And the JSON should contain:
      """
      {"status":"ok"}
      """
