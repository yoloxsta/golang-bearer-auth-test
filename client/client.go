package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// APIClient handles HTTP requests to the API
type APIClient struct {
	baseURL     string
	bearerToken string
	httpClient  *http.Client
}

// NewAPIClient creates a new API client with timeout
func NewAPIClient(baseURL, bearerToken string, timeout time.Duration) *APIClient {
	return &APIClient{
		baseURL:     baseURL,
		bearerToken: bearerToken,
		httpClient: &http.Client{
			Timeout: timeout,
		},
	}
}

// Get performs a GET request to the specified endpoint
func (c *APIClient) Get(ctx context.Context, endpoint string, result interface{}) error {
	return c.doRequest(ctx, http.MethodGet, endpoint, nil, result)
}

// Post performs a POST request to the specified endpoint
func (c *APIClient) Post(ctx context.Context, endpoint string, data interface{}, result interface{}) error {
	return c.doRequest(ctx, http.MethodPost, endpoint, data, result)
}

// Put performs a PUT request to the specified endpoint
func (c *APIClient) Put(ctx context.Context, endpoint string, data interface{}, result interface{}) error {
	return c.doRequest(ctx, http.MethodPut, endpoint, data, result)
}

// Patch performs a PATCH request to the specified endpoint
func (c *APIClient) Patch(ctx context.Context, endpoint string, data interface{}, result interface{}) error {
	return c.doRequest(ctx, http.MethodPatch, endpoint, data, result)
}

// Delete performs a DELETE request to the specified endpoint
func (c *APIClient) Delete(ctx context.Context, endpoint string) error {
	return c.doRequest(ctx, http.MethodDelete, endpoint, nil, nil)
}

// doRequest is the core method that handles all HTTP requests
func (c *APIClient) doRequest(ctx context.Context, method, endpoint string, data interface{}, result interface{}) error {
	url := c.baseURL + endpoint
	
	var bodyReader io.Reader
	if data != nil {
		jsonData, err := json.Marshal(data)
		if err != nil {
			return fmt.Errorf("failed to marshal request data: %w", err)
		}
		bodyReader = strings.NewReader(string(jsonData))
	}
	
	req, err := http.NewRequestWithContext(ctx, method, url, bodyReader)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	
	// Add Authorization header with Bearer token
	req.Header.Set("Authorization", "Bearer "+c.bearerToken)
	req.Header.Set("Content-Type", "application/json")
	
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()
	
	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}
	
	// Check status code
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("API returned status %d: %s", resp.StatusCode, string(body))
	}
	
	// Parse JSON response if result is provided
	if result != nil && len(body) > 0 {
		if err := json.Unmarshal(body, result); err != nil {
			return fmt.Errorf("failed to parse JSON response: %w", err)
		}
	}
	
	return nil
}
