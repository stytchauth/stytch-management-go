package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

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

// Get retrieves the RBAC policy for an environment.
func (c *RBACPolicyClient) Get(
	ctx context.Context,
	body rbacpolicy.GetRequest,
) (*rbacpolicy.GetResponse, error) {
	var resp rbacpolicy.GetResponse
	err := c.client.NewRequest(
		ctx,
		http.MethodGet,
		fmt.Sprintf("/pwa/v3/projects/%s/environments/%s/rbac_policy", body.Project, body.Environment),
		nil,
		nil,
		&resp)
	if err != nil {
		return nil, err
	}
	return &resp, err
}

// Set updates the RBAC policy for an environment.
func (c *RBACPolicyClient) Set(
	ctx context.Context,
	body rbacpolicy.SetRequest,
) (*rbacpolicy.SetResponse, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	var resp rbacpolicy.SetResponse
	err = c.client.NewRequest(
		ctx,
		http.MethodPut,
		fmt.Sprintf("/pwa/v3/projects/%s/environments/%s/rbac_policy", body.Project, body.Environment),
		nil,
		jsonBody,
		&resp)
	if err != nil {
		return nil, err
	}
	return &resp, err
}
