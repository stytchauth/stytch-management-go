package api

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/stytchauth/stytch-management-go/v1/pkg/api/internal"
	"github.com/stytchauth/stytch-management-go/v1/pkg/models/publictokens"
)

type PublicTokensClient struct {
	client *internal.Client
}

func newPublicTokensClient(c *internal.Client) *PublicTokensClient {
	return &PublicTokensClient{client: c}
}

// GetAll retrieves all the active public tokens defined for a project.
func (c *PublicTokensClient) GetAll(ctx context.Context, body publictokens.GetAllPublicTokensRequest) (*publictokens.GetAllPublicTokensResponse, error) {
	var resp publictokens.GetAllPublicTokensResponse
	err := c.client.NewRequest(
		ctx,
		"GET",
		fmt.Sprintf("/v1/projects/%s/public_tokens", body.ProjectID),
		nil,
		nil,
		&resp)

	return &resp, err
}

// Create creates a new public token for a project.
func (c *PublicTokensClient) Create(ctx context.Context, body publictokens.CreatePublicTokenRequest) (*publictokens.CreatePublicTokenResponse, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	var res publictokens.CreatePublicTokenResponse
	err = c.client.NewRequest(
		ctx,
		"POST",
		fmt.Sprintf("/v1/projects/%s/public_tokens", body.ProjectID),
		nil,
		jsonBody,
		&res)
	return &res, err
}

// Delete deletes a public token for a project.
func (c *PublicTokensClient) Delete(ctx context.Context, body publictokens.DeletePublicTokenRequest) (*publictokens.DeletePublicTokenResponse, error) {
	var res publictokens.DeletePublicTokenResponse
	err := c.client.NewRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/v1/projects/%s/public_tokens/%s", body.ProjectID, body.PublicToken),
		nil,
		nil,
		&res)
	return &res, err
}
