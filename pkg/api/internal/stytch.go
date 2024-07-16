package internal

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/stytchauth/stytch-management-go/v1/pkg/stytcherror"
	"github.com/stytchauth/stytch-management-go/v1/pkg/version"
)

type ClientConfig struct {
	WorkspaceKeyID     string
	WorkspaceKeySecret string
	BaseURI            string
	HTTPClient         *http.Client
}

type Client struct {
	workspaceKeyID     string
	workspaceKeySecret string

	baseURI    string
	httpClient *http.Client
}

func NewClient(c ClientConfig) *Client {
	return &Client{
		workspaceKeyID:     c.WorkspaceKeyID,
		workspaceKeySecret: c.WorkspaceKeySecret,
		baseURI:            c.BaseURI,
		httpClient:         c.HTTPClient,
	}
}

// newRequest is used by Call to generate and Do a http.Request
func (c *Client) NewRequest(
	ctx context.Context,
	method string,
	path string,
	queryParams map[string]string,
	body []byte,
	v any,
) error {
	b, err := c.RawRequest(ctx, method, path, queryParams, body)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(b, v); err != nil {
		return fmt.Errorf("error decoding http request: %w", err)
	}
	return nil
}

// RawRequest sends the request and returns the successful response body as bytes. If the response
// is an error, the response body will be parsed and returned as (nil, stytcherror.Error).
//
// Prefer using NewRequest (which unmarshals the response JSON) unless you need the actual bytes.
func (c *Client) RawRequest(
	ctx context.Context,
	method string,
	path string,
	queryParams map[string]string,
	body []byte,
) ([]byte, error) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	path = string(c.baseURI) + path

	req, err := http.NewRequestWithContext(ctx, method, path, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("error creating http request: %w", err)
	}

	// add query params
	q := req.URL.Query()
	for k, v := range queryParams {
		if v != "" {
			q.Add(k, v)
		}
	}
	req.URL.RawQuery = q.Encode()

	authToken := base64.StdEncoding.EncodeToString([]byte(c.workspaceKeyID + ":" + c.workspaceKeySecret))
	req.Header.Set("Authorization", "Basic "+authToken)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "Stytch Management Go v"+version.Version)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending http request: %w", err)
	}
	defer func() {
		res.Body.Close()
	}()

	// Successful response
	if res.StatusCode == 200 || res.StatusCode == 201 {
		return io.ReadAll(res.Body)
	}

	// Attempt to unmarshal into Stytch error format
	var stytchErr stytcherror.Error
	if err = json.NewDecoder(res.Body).Decode(&stytchErr); err != nil {
		return nil, fmt.Errorf("error decoding http request: %w", err)
	}
	stytchErr.StatusCode = res.StatusCode
	return nil, stytchErr
}
