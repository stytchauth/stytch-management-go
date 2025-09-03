package api

import (
	"context"

	"github.com/stytchauth/stytch-management-go/v3/pkg/api/internal"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/secrets"
)

type SecretsClient struct {
	client *internal.Client
}

func newSecretsClient(c *internal.Client) *SecretsClient {
	return &SecretsClient{client: c}
}

// Get retrieves a secret for a project
func (c *SecretsClient) Get(ctx context.Context, body secrets.GetSecretRequest) (*secrets.GetSecretResponse, error) {
	var resp secrets.GetSecretResponse
	err := c.client.NewRequest(
		ctx,
		"GET",
		"/pwa/v3/projects/"+body.Project+"/environments/"+body.Environment+"/secrets/"+body.SecretID,
		nil,
		nil,
		&resp,
	)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetAll retrieves all secrets for a project
func (c *SecretsClient) GetAll(ctx context.Context, body secrets.GetAllSecretsRequest) (*secrets.GetAllSecretsResponse, error) {
	var resp secrets.GetAllSecretsResponse
	err := c.client.NewRequest(
		ctx,
		"GET",
		"/pwa/v3/projects/"+body.Project+"/environments/"+body.Environment+"/secrets",
		nil,
		nil,
		&resp,
	)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Create creates a secret for a project. The response has the secret value, which
// will not be exposed in future Get requests
func (c *SecretsClient) Create(ctx context.Context, body secrets.CreateSecretRequest) (*secrets.CreateSecretResponse, error) {
	var resp secrets.CreateSecretResponse
	err := c.client.NewRequest(
		ctx,
		"POST",
		"/pwa/v3/projects/"+body.Project+"/environments/"+body.Environment+"/secrets",
		nil,
		nil,
		&resp,
	)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Delete deletes a secret for a project
func (c *SecretsClient) Delete(ctx context.Context, body secrets.DeleteSecretRequest) (*secrets.DeleteSecretResponse, error) {
	var resp secrets.DeleteSecretResponse
	err := c.client.NewRequest(
		ctx,
		"DELETE",
		"/pwa/v3/projects/"+body.Project+"/environments/"+body.Environment+"/secrets/"+body.SecretID,
		nil,
		nil,
		&resp,
	)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
