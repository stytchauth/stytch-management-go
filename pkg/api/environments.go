package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/stytchauth/stytch-management-go/v3/pkg/api/internal"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/environments"
)

type EnvironmentsClient struct {
	client *internal.Client
}

func newEnvironmentsClient(c *internal.Client) *EnvironmentsClient {
	return &EnvironmentsClient{
		client: c,
	}
}

// Create creates an environment.
func (c *EnvironmentsClient) Create(
	ctx context.Context,
	body environments.CreateRequest,
) (*environments.CreateResponse, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	var resp environments.CreateResponse
	err = c.client.NewRequest(
		ctx,
		http.MethodPost,
		fmt.Sprintf("/pwa/v3/projects/%s/environments", body.Project),
		nil,
		jsonBody,
		&resp)
	if err != nil {
		return nil, err
	}
	return &resp, err
}

// Get retrieves an environment.
func (c *EnvironmentsClient) Get(
	ctx context.Context,
	body environments.GetRequest,
) (*environments.GetResponse, error) {
	if body.Environment == "" {
		return nil, fmt.Errorf("missing environment")
	}
	var resp environments.GetResponse
	err := c.client.NewRequest(
		ctx,
		http.MethodGet,
		fmt.Sprintf("/pwa/v3/projects/%s/environments/%s", body.Project, body.Environment),
		nil,
		nil,
		&resp)
	if err != nil {
		return nil, err
	}
	return &resp, err
}

// GetAll retrieves all environments in a project.
func (c *EnvironmentsClient) GetAll(
	ctx context.Context,
	body environments.GetAllRequest,
) (*environments.GetAllResponse, error) {
	var resp environments.GetAllResponse
	err := c.client.NewRequest(
		ctx,
		http.MethodGet,
		fmt.Sprintf("/pwa/v3/projects/%s/environments", body.Project),
		nil,
		nil,
		&resp)
	if err != nil {
		return nil, err
	}
	return &resp, err
}

// Update updates the environment.
func (c *EnvironmentsClient) Update(
	ctx context.Context,
	body environments.UpdateRequest,
) (*environments.UpdateResponse, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	var resp environments.UpdateResponse
	err = c.client.NewRequest(
		ctx,
		http.MethodPatch,
		fmt.Sprintf("/pwa/v3/projects/%s/environments/%s", body.Project, body.Environment),
		nil,
		jsonBody,
		&resp)
	if err != nil {
		return nil, err
	}
	return &resp, err
}

// Delete deletes an environment.
func (c *EnvironmentsClient) Delete(
	ctx context.Context,
	body environments.DeleteRequest,
) (*environments.DeleteResponse, error) {
	var resp environments.DeleteResponse
	err := c.client.NewRequest(
		ctx,
		http.MethodDelete,
		fmt.Sprintf("/pwa/v3/projects/%s/environments/%s", body.Project, body.Environment),
		nil,
		nil,
		&resp)
	if err != nil {
		return nil, err
	}
	return &resp, err
}
