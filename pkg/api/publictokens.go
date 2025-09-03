package api

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/stytchauth/stytch-management-go/v3/pkg/api/internal"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/publictokens"
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
		fmt.Sprintf("/pwa/v3/projects/%s/environments/%s/public_tokens", body.Project, body.Environment),
		nil,
		nil,
		&resp)

	if err != nil {
		return nil, err
	}
	return &resp, nil
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
		fmt.Sprintf("/pwa/v3/projects/%s/environments/%s/public_tokens", body.Project, body.Environment),
		nil,
		jsonBody,
		&res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// Delete deletes a public token for a project.
func (c *PublicTokensClient) Delete(ctx context.Context, body publictokens.DeleteRequest) (*publictokens.DeleteResponse, error) {
	var res publictokens.DeleteResponse
	err := c.client.NewRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/pwa/v3/projects/%s/environments/%s/public_tokens/%s", body.Project, body.Environment, body.PublicToken),
		nil,
		nil,
		&res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
