package api

import (
	"context"
	"encoding/json"

	"github.com/stytchauth/stytch-management-go/v3/pkg/api/internal"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/rbacpolicy"
)

type RBACPolicyClient struct {
	client *internal.Client
}

func newRBACPolicyClient(c *internal.Client) *RBACPolicyClient {
	return &RBACPolicyClient{
		client: c,
	}
}

// Get retrieves the RBAC policy for a project
func (c *RBACPolicyClient) Get(
	ctx context.Context,
	body rbacpolicy.GetRequest,
) (*rbacpolicy.GetResponse, error) {
	var res rbacpolicy.GetResponse
	err := c.client.NewRequest(ctx, "GET", "/pwa/v3/projects/"+body.Project+"/environments/"+body.Environment+"/rbac_policy", nil, nil, &res)
	return &res, err
}

// Set updates the RBAC policy for a project
func (c *RBACPolicyClient) Set(
	ctx context.Context,
	body rbacpolicy.SetRequest,
) (*rbacpolicy.SetResponse, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	var res rbacpolicy.SetResponse
	err = c.client.NewRequest(ctx, "PUT", "/pwa/v3/projects/"+body.Project+"/environments/"+body.Environment+"/rbac_policy", nil, jsonBody, &res)
	return &res, err
}
