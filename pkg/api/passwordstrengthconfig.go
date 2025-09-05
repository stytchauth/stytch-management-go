package api

import (
	"context"
	"encoding/json"
	"fmt"
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

// Get retrieves the password strength configuration for an environment.
func (c *PasswordStrengthConfigClient) Get(
	ctx context.Context,
	body passwordstrengthconfig.GetRequest,
) (*passwordstrengthconfig.GetResponse, error) {
	var res passwordstrengthconfig.GetResponse
	err := c.client.NewRequest(
		ctx,
		http.MethodGet,
		fmt.Sprintf("/pwa/v3/projects/%s/environments/%s/password_strength_config", body.Project, body.Environment),
		nil,
		nil,
		&res,
	)
	if err != nil {
		return nil, err
	}
	return &res, err
}

// Set updates the password strength configuration for an environment.
func (c *PasswordStrengthConfigClient) Set(
	ctx context.Context,
	body passwordstrengthconfig.SetRequest,
) (*passwordstrengthconfig.SetResponse, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	var res passwordstrengthconfig.SetResponse
	err = c.client.NewRequest(
		ctx,
		http.MethodPut,
		fmt.Sprintf("/pwa/v3/projects/%s/environments/%s/password_strength_config", body.Project, body.Environment),
		nil,
		jsonBody,
		&res,
	)
	if err != nil {
		return nil, err
	}
	return &res, err
}
