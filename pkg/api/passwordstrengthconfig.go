package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/stytchauth/stytch-management-go/v3/pkg/api/internal"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/passwordstrengthconfig"
)

type PasswordStrengthConfigClient struct {
	client *internal.Client
}

func newPasswordStrengthConfigClient(c *internal.Client) *PasswordStrengthConfigClient {
	return &PasswordStrengthConfigClient{
		client: c,
	}
}

// Get retrieves the password strength configuration for a project environment
func (c *PasswordStrengthConfigClient) Get(
	ctx context.Context,
	body passwordstrengthconfig.GetRequest,
) (*passwordstrengthconfig.GetResponse, error) {
	var res passwordstrengthconfig.GetResponse
	err := c.client.NewRequest(ctx, http.MethodGet, "/pwa/v3/projects/"+body.Project+"/environments/"+body.Environment+"/password_strength_config", nil, nil, &res)
	return &res, err
}

// Set updates the password strength configuration for a project environment
func (c *PasswordStrengthConfigClient) Set(
	ctx context.Context,
	body passwordstrengthconfig.SetRequest,
) (*passwordstrengthconfig.SetResponse, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	var res passwordstrengthconfig.SetResponse
	err = c.client.NewRequest(ctx, http.MethodPut, "/pwa/v3/projects/"+body.Project+"/environments/"+body.Environment+"/password_strength_config", nil, jsonBody, &res)
	return &res, err
}
