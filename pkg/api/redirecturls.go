package api

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/stytchauth/stytch-management-go/v2/pkg/api/internal"
	"github.com/stytchauth/stytch-management-go/v2/pkg/models/redirecturls"
)

type RedirectURLsClient struct {
	client *internal.Client
}

func newRedirectURLsClient(c *internal.Client) *RedirectURLsClient {
	return &RedirectURLsClient{client: c}
}

// Create creates a redirect URL for a project
func (c *RedirectURLsClient) Create(
	ctx context.Context,
	body redirecturls.CreateRequest,
) (*redirecturls.CreateResponse, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	var res redirecturls.CreateResponse
	err = c.client.NewRequest(
		ctx,
		"POST",
		fmt.Sprintf("/v1/projects/%s/redirect_urls", body.ProjectID),
		nil,
		jsonBody,
		&res)
	return &res, err
}

// GetAll retrieves all redirect URLs for a project
func (c *RedirectURLsClient) GetAll(
	ctx context.Context,
	body redirecturls.GetAllRequest,
) (*redirecturls.GetAllResponse, error) {
	var resp redirecturls.GetAllResponse
	err := c.client.NewRequest(
		ctx,
		"GET",
		fmt.Sprintf("/v1/projects/%s/redirect_urls/all", body.ProjectID),
		nil,
		nil,
		&resp)

	return &resp, err
}

// Get retrieves a redirect URL for a project
func (c *RedirectURLsClient) Get(
	ctx context.Context,
	body redirecturls.GetRequest,
) (*redirecturls.GetResponse, error) {
	var res redirecturls.GetResponse
	err := c.client.NewRequest(
		ctx,
		"GET",
		fmt.Sprintf("/v1/projects/%s/redirect_urls?url=%s", body.ProjectID, body.URL),
		nil,
		nil,
		&res)
	return &res, err
}

// Update updates the valid types for a redirect URL for a project
func (c *RedirectURLsClient) Update(
	ctx context.Context,
	body redirecturls.UpdateRequest,
) (*redirecturls.UpdateResponse, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	var res redirecturls.UpdateResponse
	err = c.client.NewRequest(
		ctx,
		"PUT",
		fmt.Sprintf("/v1/projects/%s/redirect_urls?url=%s", body.ProjectID, body.RedirectURL.URL),
		nil,
		jsonBody,
		&res)
	return &res, err
}

// Delete deletes a redirect URL for a project
func (c *RedirectURLsClient) Delete(
	ctx context.Context,
	body redirecturls.DeleteRequest,
) (*redirecturls.DeleteResponse, error) {
	var res redirecturls.DeleteResponse
	err := c.client.NewRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/v1/projects/%s/redirect_urls?url=%s", body.ProjectID, body.URL),
		nil,
		nil,
		&res)
	return &res, err
}
