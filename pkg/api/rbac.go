package api

import (
	"context"
	"encoding/json"

	"github.com/stytchauth/stytch-management-go/v1/pkg/api/internal"
	"github.com/stytchauth/stytch-management-go/v1/pkg/models/rbac"
)

type RBACClient struct {
	client *internal.Client
}

func newRBACClient(c *internal.Client) *RBACClient {
	return &RBACClient{
		client: c,
	}
}

func (c *RBACClient) GetPolicy(
	ctx context.Context,
	body rbac.GetPolicyRequest,
) (*rbac.GetPolicyResponse, error) {
	var res rbac.GetPolicyResponse
	err := c.client.NewRequest(ctx, "GET", "/v1/projects/"+body.ProjectID+"/rbac/policy", nil, nil, &res)
	return &res, err
}

func (c *RBACClient) SetPolicy(
	ctx context.Context,
	body rbac.SetPolicyRequest,
) (*rbac.SetPolicyResponse, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	var res rbac.SetPolicyResponse
	err = c.client.NewRequest(ctx, "POST", "/v1/projects/"+body.ProjectID+"/rbac/policy", nil, jsonBody, &res)
	return &res, err
}
