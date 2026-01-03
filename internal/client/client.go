package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
}

func NewClient() *Client {
	baseURL := os.Getenv("BURNERVPN_API_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}

	return &Client{
		baseURL:    baseURL,
		httpClient: &http.Client{},
	}
}

func (c *Client) ListRegions() (*RegionsResponse, error) {
	resp, err := c.httpClient.Get(fmt.Sprintf("%s/regions", c.baseURL))
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var regions RegionsResponse
	err = json.NewDecoder(resp.Body).Decode(&regions)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &regions, nil
}

func (c *Client) CreateSession(region string) (*CreateSessionResponse, error) {
	req := CreateSessionRequest{
		Region: region,
	}
	bodyBytes, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := c.httpClient.Post(fmt.Sprintf("%s/sessions", c.baseURL),
		"application/json",
		bytes.NewReader(bodyBytes),
	)

	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result CreateSessionResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	return &result, nil
}

func (c *Client) DeleteSession(sessionID string) (*DeleteSessionResponse, error) {
	url := fmt.Sprintf("%s/sessions/%s", c.baseURL, sessionID)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result DeleteSessionResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	return &result, nil
}

func (c *Client) GetUsage(sessionID string) (*UsageResponse, error) {
	url := fmt.Sprintf("%s/sessions/%s/usage", c.baseURL, sessionID)
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result UsageResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	return &result, nil
}
