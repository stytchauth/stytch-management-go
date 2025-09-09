package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/stytchauth/stytch-management-go/v3/pkg/api/internal"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/publictokens"
)

type PublicTokensClient struct {
	client *internal.Client
}

func newPublicTokensClient(c *internal.Client) *PublicTokensClient {
	return &PublicTokensClient{client: c}
}

// Create creates a new public token for an environment.
func (c *PublicTokensClient) Create(
	ctx context.Context,
	body publictokens.CreateRequest,
) (*publictokens.CreateResponse, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	var res publictokens.CreateResponse
	err = c.client.NewRequest(
		ctx,
		http.MethodPost,
		fmt.Sprintf("/pwa/v3/projects/%s/environments/%s/public_tokens", body.Project, body.Environment),
		nil,
		jsonBody,
		&res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// Get retrieves a public token for an environment.
func (c *PublicTokensClient) Get(
	ctx context.Context,
	body publictokens.GetRequest,
) (*publictokens.GetResponse, error) {
	if body.PublicToken == "" {
		return nil, fmt.Errorf("missing public token")
	}
	var resp publictokens.GetResponse
	err := c.client.NewRequest(
		ctx,
		http.MethodGet,
		fmt.Sprintf("/pwa/v3/projects/%s/environments/%s/public_tokens/%s", body.Project, body.Environment, body.PublicToken),
		nil,
		nil,
		&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetAll retrieves all the active public tokens defined for an environment.
func (c *PublicTokensClient) GetAll(
	ctx context.Context,
	body publictokens.GetAllRequest,
) (*publictokens.GetAllResponse, error) {
	var resp publictokens.GetAllResponse
	err := c.client.NewRequest(
		ctx,
		http.MethodGet,
		fmt.Sprintf("/pwa/v3/projects/%s/environments/%s/public_tokens", body.Project, body.Environment),
		nil,
		nil,
		&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Delete deletes a public token for an environment.
func (c *PublicTokensClient) Delete(
	ctx context.Context,
	body publictokens.DeleteRequest,
) (*publictokens.DeleteResponse, error) {
	var res publictokens.DeleteResponse
	err := c.client.NewRequest(
		ctx,
		http.MethodDelete,
		fmt.Sprintf("/pwa/v3/projects/%s/environments/%s/public_tokens/%s", body.Project, body.Environment, body.PublicToken),
		nil,
		nil,
		&res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
