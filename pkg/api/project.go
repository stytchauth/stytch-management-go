package api

import (
	"context"
	"encoding/json"

	"github.com/stytchauth/stytch-management-go/v3/pkg/api/internal"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/project"
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
	body project.CreateRequest,
) (*project.CreateResponse, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	var res project.CreateResponse
	err = c.client.NewRequest(ctx, "POST", "/pwa/v3/projects", nil, jsonBody, &res)
	if err != nil {
		return nil, err
	}
	return &res, err
}

// Get retrieves a project.
func (c *ProjectClient) Get(
	ctx context.Context,
	body project.GetRequest,
) (*project.GetResponse, error) {
	var res project.GetResponse
	err := c.client.NewRequest(ctx, "GET", "/pwa/v3/projects/"+body.Project, nil, nil, &res)
	if err != nil {
		return nil, err
	}
	return &res, err
}

// GetAll retrieves all projects in a workspace.
func (c *ProjectClient) GetAll(
	ctx context.Context,
	body project.GetAllRequest,
) (*project.GetAllResponse, error) {
	var res project.GetAllResponse
	err := c.client.NewRequest(ctx, "GET", "/pwa/v3/projects", nil, nil, &res)
	if err != nil {
		return nil, err
	}
	return &res, err
}

// Update updates the project.
func (c *ProjectClient) Update(
	ctx context.Context,
	body project.UpdateRequest,
) (*project.UpdateResponse, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	var res project.UpdateResponse
	err = c.client.NewRequest(ctx, "PATCH", "/pwa/v3/projects/"+body.Project, nil, jsonBody, &res)
	if err != nil {
		return nil, err
	}
	return &res, err
}

// Delete deletes a project and all of its environments.
func (c *ProjectClient) Delete(
	ctx context.Context,
	body project.DeleteRequest,
) (*project.DeleteResponse, error) {
	var res project.DeleteResponse
	err := c.client.NewRequest(ctx, "DELETE", "/pwa/v3/projects/"+body.Project, nil, nil, &res)
	if err != nil {
		return nil, err
	}
	return &res, err
}
