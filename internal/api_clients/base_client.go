package api_clients

import (
	"bytes"
	"computer-manager/internal/config"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type ApiClients struct {
	AdminAlarm *AdminAlarmClient
}

func NewApiClients(cfg config.ThirdPartyServicesConfig) *ApiClients {
	return &ApiClients{
		AdminAlarm: NewAdminAlarmClient(cfg.AdminAlarm.BaseURL),
	}
}

type BaseClient struct {
	BaseURL    string
	httpClient *http.Client
}

func NewBaseClient(baseURL string) BaseClient {
	c := BaseClient{
		BaseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}

	return c
}

func (c *BaseClient) Request(ctx context.Context, method, path string, body any, response any, headers map[string]string) error {
	url := fmt.Sprintf("%s%s", c.BaseURL, path)
	var req *http.Request
	var err error

	switch v := body.(type) {
	case *bytes.Buffer:
		req, err = http.NewRequestWithContext(ctx, method, url, v)
		if err != nil {
			return fmt.Errorf("failed to create request with multipart form data: %w", err)
		}
	default:
		var bodyReader io.Reader
		if body != nil {
			jsonBody, err := json.Marshal(body)
			if err != nil {
				return fmt.Errorf("failed to marshal request body: %w", err)
			}
			bodyReader = bytes.NewReader(jsonBody)
			if headers != nil {
				if _, exists := headers["Content-Type"]; !exists {
					headers["Content-Type"] = "application/json"
				}
			}
		}
		req, err = http.NewRequestWithContext(ctx, method, url, bodyReader)
		if err != nil {
			return fmt.Errorf("failed to create request: %w", err)
		}
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("request failed with status %d: %s", resp.StatusCode, string(body))
	}

	if response != nil {
		switch response.(type) {
		case *[]byte:
			data, err := io.ReadAll(resp.Body)
			if err != nil {
				return fmt.Errorf("failed to read binary response: %w", err)
			}
			fmt.Printf("Binary response: %s\n", data)
		default:
			if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
				return fmt.Errorf("failed to decode response: %w", err)
			}
		}
	}

	return nil
}
