package api

import (
	"context"
	"encoding/json"

	"github.com/stytchauth/stytch-management-go/v3/pkg/api/internal"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/projects"
)

type ProjectClient struct {
	client *internal.Client
}

func newProjectClient(c *internal.Client) *ProjectClient {
	return &ProjectClient{
		client: c,
	}
}

// Create creates a project, including both a live and test environment.
func (c *ProjectClient) Create(
	ctx context.Context,
	body projects.CreateRequest,
) (*projects.CreateResponse, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	var res projects.CreateResponse
	err = c.client.NewRequest(ctx, "POST", "/pwa/v3/projects", nil, jsonBody, &res)
	if err != nil {
		return nil, err
	}
	return &res, err
}

// Get retrieves a project.
func (c *ProjectClient) Get(
	ctx context.Context,
	body projects.GetRequest,
) (*projects.GetResponse, error) {
	var res projects.GetResponse
	err := c.client.NewRequest(ctx, "GET", "/pwa/v3/projects/"+body.Project, nil, nil, &res)
	if err != nil {
		return nil, err
	}
	return &res, err
}

// GetAll retrieves all projects in a workspace.
func (c *ProjectClient) GetAll(
	ctx context.Context,
	body projects.GetAllRequest,
) (*projects.GetAllResponse, error) {
	var res projects.GetAllResponse
	err := c.client.NewRequest(ctx, "GET", "/pwa/v3/projects", nil, nil, &res)
	if err != nil {
		return nil, err
	}
	return &res, err
}

// Update updates the project.
func (c *ProjectClient) Update(
	ctx context.Context,
	body projects.UpdateRequest,
) (*projects.UpdateResponse, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	var res projects.UpdateResponse
	err = c.client.NewRequest(ctx, "PATCH", "/pwa/v3/projects/"+body.Project, nil, jsonBody, &res)
	if err != nil {
		return nil, err
	}
	return &res, err
}

// Delete deletes a project and all of its environments.
func (c *ProjectClient) Delete(
	ctx context.Context,
	body projects.DeleteRequest,
) (*projects.DeleteResponse, error) {
	var res projects.DeleteResponse
	err := c.client.NewRequest(ctx, "DELETE", "/pwa/v3/projects/"+body.Project, nil, nil, &res)
	if err != nil {
		return nil, err
	}
	return &res, err
}
