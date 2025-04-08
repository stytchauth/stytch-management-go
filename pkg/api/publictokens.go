package api

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/stytchauth/stytch-management-go/v2/pkg/api/internal"
	"github.com/stytchauth/stytch-management-go/v2/pkg/models/publictokens"
)

type PublicTokensClient struct {
	client *internal.Client
}

func newPublicTokensClient(c *internal.Client) *PublicTokensClient {
	return &PublicTokensClient{client: c}
}

// GetAll retrieves all the active public tokens defined for a project.
func (c *PublicTokensClient) GetAll(ctx context.Context, body publictokens.GetAllRequest) (*publictokens.GetAllResponse, error) {
	var resp publictokens.GetAllResponse
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
func (c *PublicTokensClient) Create(ctx context.Context, body publictokens.CreateRequest) (*publictokens.CreateResponse, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	var res publictokens.CreateResponse
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
func (c *PublicTokensClient) Delete(ctx context.Context, body publictokens.DeleteRequest) (*publictokens.DeleteResponse, error) {
	var res publictokens.DeleteResponse
	err := c.client.NewRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/v1/projects/%s/public_tokens/%s", body.ProjectID, body.PublicToken),
		nil,
		nil,
		&res)
	return &res, err
}
