package api

import (
	"context"
	"encoding/json"

	"github.com/stytchauth/stytch-management-go/v1/pkg/api/internal"
	"github.com/stytchauth/stytch-management-go/v1/pkg/models/passwordstrengthconfig"
)

type PasswordStrengthConfigClient struct {
	client *internal.Client
}

func newPasswordStrengthConfigClient(c *internal.Client) *PasswordStrengthConfigClient {
	return &PasswordStrengthConfigClient{
		client: c,
	}
}

// Get retrieves the password strength configuration for a project
func (c *PasswordStrengthConfigClient) Get(
	ctx context.Context,
	body passwordstrengthconfig.GetRequest,
) (*passwordstrengthconfig.GetResponse, error) {
	var res passwordstrengthconfig.GetResponse
	err := c.client.NewRequest(ctx, "GET", "/v1/projects/"+body.ProjectID+"/password_strength", nil, nil, &res)
	return &res, err
}

// Set updates the password strength configuration for a project
func (c *PasswordStrengthConfigClient) Set(
	ctx context.Context,
	body passwordstrengthconfig.SetRequest,
) (*passwordstrengthconfig.SetResponse, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	var res passwordstrengthconfig.SetResponse
	err = c.client.NewRequest(ctx, "PUT", "/v1/projects/"+body.ProjectID+"/password_strength", nil, jsonBody, &res)
	return &res, err
}
