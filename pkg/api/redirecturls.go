package api

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/stytchauth/stytch-management-go/v1/pkg/api/internal"
	"github.com/stytchauth/stytch-management-go/v1/pkg/models/redirecturls"
)

type RedirectURLsClient struct {
	client *internal.Client
}

func newRedirectURLsClient(c *internal.Client) *RedirectURLsClient {
	return &RedirectURLsClient{client: c}
}

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

func (c *RedirectURLsClient) GetAll(
	ctx context.Context,
	body redirecturls.GetAllRequest,
) (*redirecturls.GetAllResponse, error) {
	var resp redirecturls.GetAllResponse
	err := c.client.NewRequest(
		ctx,
		"GET",
		fmt.Sprintf("/v1/projects/%s/redirect_urls", body.ProjectID),
		nil,
		nil,
		&resp)

	return &resp, err
}

func (c *RedirectURLsClient) RemoveValidType(
	ctx context.Context,
	body redirecturls.RemoveValidTypeRequest,
) (*redirecturls.RemoveValidTypeResponse, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	var res redirecturls.RemoveValidTypeResponse
	err = c.client.NewRequest(
		ctx,
		"POST",
		fmt.Sprintf("/v1/projects/%s/redirect_urls/remove_valid_type", body.ProjectID),
		nil,
		jsonBody,
		&res)
	return &res, err
}
